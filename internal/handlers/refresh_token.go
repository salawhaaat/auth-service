package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/salawhaaat/auth-service/internal/services"
)

func HandleRefreshToken(c *gin.Context) {
	// Get tokens from cookies
	oldAccessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(400, gin.H{"error": "Access token not found"})
		return
	}

	oldRefreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(400, gin.H{"error": "Refresh token not found"})
		return
	}

	newTokens, err := services.RefreshToken(oldRefreshToken, oldAccessToken)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Set new tokens in cookies
	c.SetCookie("access_token", newTokens.AccessToken, AccessTokenMaxAge, "/", "", isSecure, isHttpOnly)
	c.SetCookie("refresh_token", newTokens.RefreshToken, RefreshTokenMaxAge, "/", "", isSecure, isHttpOnly)

	c.JSON(200, gin.H{"message": "Tokens are refreshed and set in the cookies"})
}
