package controller

import (
	"avatar/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNotificationRequest struct {
	Key string `json:"key" binding:"required"`
}

func GetNotification(c *gin.Context) {
	// Get the key from the request body
	var param GetNotificationRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	fmt.Println("Notification:", param.Key)
	// Call the service to get the notification
	notification, err := service.GetNotification(param.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy vấn"})
		return
	}

	// Return the notification as JSON response
	c.JSON(http.StatusOK, gin.H{"notification": notification})
}
