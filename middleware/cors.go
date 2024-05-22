package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureCORS(c *gin.Context) {
	//TODO ALL ROUTES ARE ALLOWED!!!
	origin := c.GetHeader("Origin")

	if origin == "http://localhost:3000" {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
