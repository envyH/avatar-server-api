package main

import (
	"avatar/config"
	"avatar/internal/db"
	"avatar/internal/router"
	"fmt"
	"os"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	db.InitDB(cfg.DBURL)
	defer db.CloseDB()

	// Setup router
	r := router.SetupRouter(cfg.GIN_MODE)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Server chạy trên cổng %s\n", port)
	r.Run(":" + port) // Gin's built-in method to start the server
}
