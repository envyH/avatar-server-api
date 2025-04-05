package router

import (
	"avatar/internal/router/farm"
	"avatar/internal/router/global"
	"avatar/internal/router/notification"
	"avatar/internal/router/player"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

// SetupRouter initializes and returns the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define a custom function to safely render HTML
	r.SetFuncMap(template.FuncMap{
		"safeHTML": func(content string) template.HTML {
			return template.HTML(content)
		},
	})

	// Serve static files
	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/icon.png")

	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Route
	r.GET("/", func(c *gin.Context) {
		msg := "Avatar 258 Mod by Envy"
		// Đọc file README và CHANGELOG
		readmeBytes, _ := os.ReadFile("docs/README.md")
		changelogBytes, _ := os.ReadFile("docs/CHANGELOG.md")
		// Convert Markdown -> HTML
		readmeContent := blackfriday.Run(readmeBytes, blackfriday.WithNoExtensions())
		changelogContent := blackfriday.Run(changelogBytes)
		c.HTML(200, "index.html", gin.H{
			"title":     msg,
			"message":   msg,
			"readme":    string(readmeContent),
			"changelog": string(changelogContent),
		})
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

		player.RegisterPlayerRoutes(api)
		notification.NotificationRoutes(api)
		farm.FarmRoutes(api)
		global.GlobalRoutes(api)
	}

	return r
}
