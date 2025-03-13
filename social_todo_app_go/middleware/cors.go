package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func AllowCors() func(ctx *gin.Context) {
	return func(context *gin.Context) {
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:5173"}, // Allow Vue app origin
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"cache-control", "expires", "pragma"},
			Debug:          true, // Set to false in production
		})
		c.HandlerFunc(context.Writer, context.Request)
		context.Next()
	}
}
