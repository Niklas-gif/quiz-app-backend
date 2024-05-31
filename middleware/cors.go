package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureCORS(c *gin.Context) {
	origin := c.GetHeader("Origin")

	allowedOrigins := []string{
		"http://localhost:3000",
		//TODO setup env for local debug server and add the allowed origins there
	}

	isAllowed := false
	for _, o := range allowedOrigins {
		if origin == o {
			isAllowed = true
			break
		}
	}

	if isAllowed {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
	}

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
