package farm

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

func FarmRoutes(r *gin.RouterGroup) {
	r.GET("/farm-data", controller.GetFarmData)
}
