package controllers

import (
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
)

// CreateLabel godoc
// @Summary Create Label data
// @Description Add by JSON Label
// @Tags label
// @Accept  json
// @Produce  json
// @Param label body models.Label true "Create Label"
// @success 200 {object} models.Label
// @Router /label [post]

func CreateLabel(c *gin.Context) {
	var label models.Label
	c.ShouldBindJSON(&label)
	if err := config.DB.Create(&label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Label Cannot Be Created!"})
		return
	}
	c.JSON(http.StatusOK, label)
}
