package global

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

func GlobalRoutes(r *gin.RouterGroup) {
	r.POST("/image/b2i", controller.Byte2Image)
	r.POST("/captcha/train", controller.TrainCaptcha)
}
