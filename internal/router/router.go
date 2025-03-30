package router

import (
	"avatar/internal/router/farm"
	"avatar/internal/router/notification"
	"avatar/internal/router/player"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func getTemplatePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(dir, "templates", "*")
}

// SetupRouter initializes and returns the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob(getTemplatePath())

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
		// Register player routes
		player.RegisterPlayerRoutes(api)
		notification.NotificationRoutes(api)
		farm.FarmRoutes(api)
	}

	return r
}
