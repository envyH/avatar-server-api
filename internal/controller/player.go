package controller

import (
	"avatar/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPlayers handles the request to retrieve all players
func GetPlayers(c *gin.Context) {
	players, err := service.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy vấn"})
		return
	}

	c.JSON(http.StatusOK, players)
}

// UpdateScore handles the request to update a player's score
func UpdateScore(c *gin.Context) {
	var data struct {
		ID    int `json:"id"`
		Score int `json:"score"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	if err := service.UpdatePlayerScore(data.ID, data.Score); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi cập nhật điểm"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật điểm thành công"})
}
