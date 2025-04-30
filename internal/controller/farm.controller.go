package controller

import (
	"avatar/internal/service"
	"encoding/json"
	"fmt"
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

type SyncAnimalRequest struct {
	ID         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Desc       string `json:"desc"`
	BornTime   int    `json:"born_time"`
	MatureTime int    `json:"mature_time"`
}

func SyncAnimal(c *gin.Context) {
	var param SyncAnimalRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	jsonBytes, _ := json.MarshalIndent(param, "", "  ")
	fmt.Println("Received animal JSON:\n", string(jsonBytes))
	c.JSON(http.StatusOK, gin.H{"message": "Sync animal success"})
}
