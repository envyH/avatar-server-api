package controller

import (
	"avatar/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFarmData(c *gin.Context) {
	farmData, err := service.GetFarmData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy dữ liệu trang trại"})
		return
	}

	c.JSON(http.StatusOK, farmData)
}
