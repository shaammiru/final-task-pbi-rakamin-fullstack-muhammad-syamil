package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/helpers"
	"net/http"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"data":    nil,
				"error":   "Token is required",
			})
			c.Abort()
			return
		}

		claims, err := helpers.GetTokenClaims(token[7:])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"data":    nil,
				"error":   "Invalid (malformed or expired) token",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
