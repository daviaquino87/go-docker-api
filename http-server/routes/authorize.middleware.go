package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
			os.Exit(1) 
		}

		fixedToken := os.Getenv("TOKEN")

		if fixedToken == "" {
			fmt.Println("TOKEN variable is required in the env file")
			os.Exit(1) 
			
			return
		}

		expectedToken := "Bearer " + fixedToken

		if token != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}