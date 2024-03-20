package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/salawhaaat/auth-service/internal/models"
	"github.com/salawhaaat/auth-service/internal/services"
)

const (
	AccessTokenMaxAge  = 3600  // 1 hour
	RefreshTokenMaxAge = 86400 // 24 hours
	isHttpOnly         = true
	isSecure           = true
)

func HandleGetToken(c *gin.Context) {
	userId := c.Param("userId")
	// Check if userId is a valid UUID
	if _, err := uuid.Parse(userId); err != nil {
		// If not, return a 400 Bad Request response
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	var token models.AuthToken
	var err error
	token, err = services.GenerateTokens(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Set cookies
	c.SetCookie("access_token", token.AccessToken, AccessTokenMaxAge, "/", "", isSecure, isHttpOnly)
	c.SetCookie("refresh_token", token.RefreshToken, RefreshTokenMaxAge, "/", "", isSecure, isHttpOnly)

	c.JSON(200, gin.H{"message": "Tokens are set in the cookies"})
}
