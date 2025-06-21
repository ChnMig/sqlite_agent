package middleware

import (
	"github.com/gin-gonic/gin"

	"sqlite-agent/api/response"
	"sqlite-agent/config"
)

// AuthMiddleware checks if the request contains a valid auth header.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("auth")
		if authHeader == "" || authHeader != config.ApiAuth {
			response.ReturnError(c, response.UNAUTHENTICATED, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
