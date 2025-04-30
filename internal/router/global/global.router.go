package global

import (
	"avatar/internal/controller"

	"github.com/gin-gonic/gin"
)

func GlobalRoutes(r *gin.RouterGroup) {
	r.POST("/image/b2i", controller.Byte2Image)
	r.POST("/captcha/train", controller.TrainCaptcha)
	r.POST("/dialog/input/answer", controller.GetAnswerInputDlg)
	r.POST("/crawl/icon", controller.CrawlIcon)
	r.POST("/crawl/image", controller.CrawlImage)
}
