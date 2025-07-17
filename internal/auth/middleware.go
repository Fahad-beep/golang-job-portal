package auth

import (
	"fmt"
	"job_portal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("is authMiddleware")
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "no token given"})
			c.Abort()
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
		}
		fmt.Println(claims)

		c.Set("userID", claims.UserID)
		c.Set("isAdmin", claims.IsAdmin)
		c.Next()
	}

}
