package controller

import (
	"avatar/internal/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrawlImageIconRequest struct {
	Uid  int32   `json:"uid" binding:"required"`
	ID   int16   `json:"id" binding:"required"`
	Data []uint8 `json:"data" binding:"required"`
}

func CrawlIcon(c *gin.Context) {
	var param CrawlImageIconRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	filePath, isNew, err := service.SaveIcon(param.Uid, param.ID, param.Data)
	if err != nil {
		fmt.Println("Lỗi khi lưu icon:", err)
	} else {
		if isNew {
			fmt.Println("icon mới đã được lưu tại:", filePath)
		} else {
			// fmt.Println("icon đã tồn tại và giống nhau, không ghi đè:", filePath)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "send success"})
}

func CrawlImage(c *gin.Context) {
	var param CrawlImageIconRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	filePath, isNew, err := service.SaveImageIfNotExists(param.Uid, param.ID, param.Data)
	if err != nil {
		fmt.Println("Lỗi khi lưu ảnh:", err)
	} else {
		if isNew {
			fmt.Println("Ảnh mới đã được lưu tại:", filePath)
		} else {
			// fmt.Println("Ảnh đã tồn tại và giống nhau, không ghi đè:", filePath)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "send success"})
}

type InputDlgRequest struct {
	Uid     int32  `json:"uid" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func GetAnswerInputDlg(c *gin.Context) {
	var param InputDlgRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	jsonBytes, _ := json.MarshalIndent(param, "", "  ")
	fmt.Println("Received animal JSON:\n", string(jsonBytes))
	fmt.Println("message:", param.Message)
	answer, err := service.GenerateText(param.Uid, param.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate text"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": answer})
}

type Byte2ImageRequest struct {
	Uid  int32   `json:"uid" binding:"required"`
	Id   int32   `json:"id" binding:"required"`
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
	filePathImg, _, err := service.TrainCaptcha(param.Uid, param.Data, param.Captcha)
	if err != nil {
		fmt.Println("Lỗi khi lưu file:", err)
	} else {
		fmt.Println("Ảnh đã được lưu tại:", filePathImg)
		// fmt.Println("Captcha đã được lưu tại:", filePathCaptcha)
	}
	c.JSON(http.StatusOK, gin.H{"message": "send success"})
}

func Byte2Image(c *gin.Context) {
	var param Byte2ImageRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	filePath, err := service.SaveImageCrawl(param.Uid, param.Id, param.Data)
	if err != nil {
		fmt.Println("Lỗi khi lưu ảnh:", err)
	} else {
		fmt.Println("Ảnh đã được lưu tại:", filePath)
	}
	c.JSON(http.StatusOK, gin.H{"message": "send success"})
}

type CreateQRcodeRequest struct {
	Width  int    `json:"width" binding:"required"`
	Height int    `json:"height" binding:"required"`
	Url    string `json:"url" binding:"required"`
}

func CreateQRcode(c *gin.Context) {
	var param CreateQRcodeRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, dataByte, err := service.CreateQRcode(param.Url, param.Width, param.Height)
	if err != nil {
		fmt.Println("ERROR CreateQRcode", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": dataByte})
}
