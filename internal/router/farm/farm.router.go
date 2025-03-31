package farm

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

func FarmRoutes(r *gin.RouterGroup) {
	r.POST("/farm-data", controller.GetFarmData)
}
