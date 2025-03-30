package router

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		// Render the "index.html" file
		c.HTML(200, "index.html", nil)
	})
	// Group API routes under /api/v1
	api := r.Group("/api/v1")
	{
		api.POST("/ping", func(c *gin.Context) {
			message := c.DefaultQuery("message", "pong")
			c.JSON(200, gin.H{
				"message": message,
			})
		})
		api.GET("/players", controller.GetPlayers)
		api.POST("/update-score", controller.UpdateScore)
	}

	return r
}
