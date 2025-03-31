package controller

import (
	"avatar/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Byte2ImageRequest struct {
	Uid  int32   `json:"uid" binding:"required"`
	Data []uint8 `json:"data" binding:"required"`
}

type TrainCaptchaRequest struct {
	Uid     int32   `json:"uid" binding:"required"`
	Data    []uint8 `json:"data" binding:"required"`
	Captcha string  `json:"captcha" binding:"required"`
}

func TrainCaptcha(c *gin.Context) {
	var param TrainCaptchaRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	filePathImg, filePathCaptcha, err := service.TrainCaptcha(param.Uid, param.Data, param.Captcha)
	if err != nil {
		fmt.Println("Lỗi khi lưu file:", err)
	} else {
		fmt.Println("Ảnh đã được lưu tại:", filePathImg)
		fmt.Println("Captcha đã được lưu tại:", filePathCaptcha)
	}
	c.JSON(http.StatusOK, gin.H{"message": "send success"})
}

func Byte2Image(c *gin.Context) {
	var param Byte2ImageRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	filePath, err := service.SaveImage(param.Uid, param.Data)
	if err != nil {
		fmt.Println("Lỗi khi lưu ảnh:", err)
	} else {
		fmt.Println("Ảnh đã được lưu tại:", filePath)
	}
	c.JSON(http.StatusOK, gin.H{"message": "send success"})
}
