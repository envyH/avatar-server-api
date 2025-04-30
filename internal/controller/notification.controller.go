package controller

import (
	"avatar/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNotificationRequest struct {
	Key string `json:"key" binding:"required"`
}

func GetNotification(c *gin.Context) {
	var param GetNotificationRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	notification, err := service.GetNotification(param.Key, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy vấn: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notification": notification})
}
