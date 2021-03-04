package controllers

import (
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
)

//CreateProject godoc
//@Summary Create Project data
//@Description Add by JSON Project
//@Tags project
//@Accept  json
//@Produce  json
//@Param project body models.Project true "Add Project"
//@success 200 {object} models.Project
//@Router /project [post]
func CreateProject(c *gin.Context) {
	var project models.Project
	c.ShouldBindJSON(&project)
	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project Cannot Be Created!"})
		return
	}
	c.JSON(http.StatusOK, project)
}
