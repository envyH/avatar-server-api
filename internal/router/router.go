package router

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/players", controller.GetPlayers)
	r.POST("/update-score", controller.UpdateScore)
	return r
}
