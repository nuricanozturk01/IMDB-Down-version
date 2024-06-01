package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CorsPolicy is a middleware function that sets the CORS policy for the application. (Allow all origins)
func CorsPolicy() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}
	return cors.New(corsConfig)
}
