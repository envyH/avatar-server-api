package player

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

// RegisterPlayerRoutes registers routes related to players
func RegisterPlayerRoutes(r *gin.RouterGroup) {
	r.GET("/players", controller.GetPlayers)
	r.POST("/update-score", controller.UpdateScore)
}
