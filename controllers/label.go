package controllers

import (
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func CreateLabel(c *gin.Context) {
	var label models.Label
	c.ShouldBindJSON(&label)
	if err := config.DB.Create(&label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Label Cannot Be Created!"})
		return
	}
	c.JSON(http.StatusOK, label)
}
