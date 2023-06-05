package middleware

import (
	"log"
	"miniProject2/pkg/tokenize"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		token := tokenize.SplitBearer(tokenString)
		_, err := tokenize.VerifyAccessToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		accountData, err := tokenize.GetDataUserFromToken(token)
		if err != nil {
			log.Println(err)
		}
		if !accountData.IsVerified || !accountData.IsActive {
			c.JSON(401, gin.H{"error": "User admin must be verified and active"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		token := tokenize.SplitBearer(tokenString)
		_, err := tokenize.VerifyAccessToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		accountData, err := tokenize.GetDataUserFromToken(token)
		if err != nil {
			log.Println(err)
		}
		if !accountData.IsVerified || !accountData.IsActive {
			c.JSON(401, gin.H{"error": "User admin must be verified and active"})
			c.Abort()
			return
		}
		if accountData.RoleID != 2 {
			c.JSON(401, gin.H{"error": "Only for super admin role"})
			c.Abort()
		}
		c.Next()
	}
}
