package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware returns a CORS middleware using gin-contrib/cors.
// Allowed origins are configured via the ALLOWED_ORIGINS environment variable.
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	return cors.New(config)
}

