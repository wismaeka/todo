package controllers

import (
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Created"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetTaskById(c *gin.Context) {
	task_id := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Where("task_id", task_id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetAllTask(c *gin.Context) {
	var task models.Task
	if err := config.DB.Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	task_id := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Where("task_id", task_id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.ShouldBindJSON(&task)
	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Updated!"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	task_id := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Where("task_id=?", task_id).Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task_id" + task_id: "is deleted"})
}
