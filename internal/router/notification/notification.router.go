package notification

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.RouterGroup) {
	r.POST("/notification", controller.GetNotification)
}
