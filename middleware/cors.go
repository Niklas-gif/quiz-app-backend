package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureCORS(context *gin.Context) {
	context.Header(
		"Access-Control-Allow-Origin",
		"http://localhost:3000",
	)

	context.Header(
		"Access-Control-Allow-Methods",
		"GET, POST, PUT, DELETE, OPTIONS",
	)

	if context.Request.Method == "OPTIONS" {
		context.Status(http.StatusNoContent)
	}

	context.Next()
}
