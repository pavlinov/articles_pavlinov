package middleware

import (
	"articles_pavlinov/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	log.Println("AlexeyAAAAAAAAAAAA")
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Println("Authorization header:", authHeader)

		if authHeader == "" {
			log.Println("Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if the token is a Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Println("Authorization header is not a Bearer token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with Bearer"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println("Extracted token:", token)

		userID, err := utils.ParseToken(token)
		if err != nil {
			log.Println("Error parsing token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Println("Authenticated user ID:", userID)
		c.Set("user_id", userID)
		c.Next()
	}
}
