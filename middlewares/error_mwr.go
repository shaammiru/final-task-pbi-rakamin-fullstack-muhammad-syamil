package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
				})
			}
		}()

		c.Next()
	}
}
