// routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vivekjha1213/Limit-Warden-Go/middleware"
	clientController "github.com/vivekjha1213/Limit-Warden-Go/controllers"
)

// LoadRoutes loads all routes for the application.
func LoadRoutes(r *gin.Engine) {
	// Load routes for v1 version
	loadV1Routes(r)
}

func loadV1Routes(r *gin.Engine) {
	// Group for /v1/token endpoint
	v1 := r.Group("/v1")
	{
		client := v1.Group("/token")
		{
			client.GET("/", clientController.GenerateClientKey)
		}

		// Group for /v1/ping endpoint with RateLimit middleware
		endpoint := v1.Group("/ping")
		{
			endpoint.Use(middleware.RateLimit)
			endpoint.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{"message": "Success Ok"})
			})
		}
	
	}
}