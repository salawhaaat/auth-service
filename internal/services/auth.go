package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/salawhaaat/auth-service/internal/database"
	"github.com/salawhaaat/auth-service/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var repo database.Repository

func init() {
	repo = database.NewMongoRepository()
	repo.Connect()
}

// GenerateTokens generates an access token and a refresh token
func GenerateTokens(userId string) (token models.AuthToken, err error) {
	token.AccessToken, err = generateAccessToken(userId)
	if err != nil {
		return
	}
	var refreshToken string
	refreshToken, err = generateRefreshToken()
	if err != nil {
		return
	}
	rt, err := hashRefreshToken(refreshToken)
	if err != nil {
		return
	}
	// Store the refresh token in the database
	if err := repo.StoreRefreshToken(userId, rt); err != nil {
		return models.AuthToken{}, err
	}

	token.RefreshToken = refreshToken
	return
}

// RefreshToken checks if the refresh token and pairs with access token is valid and generates new tokens
func RefreshToken(oldRefreshToken string, accessToken string) (models.AuthToken, error) {
	userId, err := getUserIdFromToken(accessToken) // get the user id from the access token claims
	if err != nil {
		return models.AuthToken{}, fmt.Errorf("invalid access token")
	}

	if !VerifyRefreshToken(userId, oldRefreshToken) {
		return models.AuthToken{}, fmt.Errorf("invalid refresh token")
	}
	return GenerateTokens(userId)
}

// VerifyRefreshToken verifies the refresh token
func VerifyRefreshToken(userId string, oldRefreshToken string) bool {
	hashedRTFromDB := repo.GetRefreshToken(userId)
	err := bcrypt.CompareHashAndPassword([]byte(hashedRTFromDB), []byte(oldRefreshToken))
	if err != nil {
		return false
	}
	return true
}

// getUserIdFromToken extracts the user id from the access token
func getUserIdFromToken(accessToken string) (string, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["jwt_id"].(string)
		return userId, nil
	}
	return "", err
}

// generateAccessToken generates an access token
func generateAccessToken(user_id string) (string, error) {
	atClaims := jwt.MapClaims{
		"jwt_id": user_id,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	return at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

// generateRefreshToken generates a refresh token
func generateRefreshToken() (string, error) {
	rt := make([]byte, 32)
	_, err := rand.Read(rt)
	return base64.StdEncoding.EncodeToString(rt), err
}

// hashRefreshToken hashes the refresh token
func hashRefreshToken(refreshToken string) (string, error) {
	hashedRT, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedRT), nil
}
