package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salawhaaat/auth-service/internal/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	r := gin.Default()
	r.GET("/auth/:userId", handlers.HandleGetToken)
	r.GET("/refresh", handlers.HandleRefreshToken)
	r.Run(":8080")
}
