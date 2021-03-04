package controllers

import (
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreateTask godoc
//@Summary Create task
//@Description Add by JSON Task
//@Tags task
//@Accept  json
//@Produce  json
//@Param task body models.Task true "Add Task"
//@success 200 {object} models.Task
//@Router /task [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)
	var project models.Project
	if err := config.DB.Find(&project).Where("project_title", task.Project.ProjectTitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project Cannot Be Found"})
		return
	}
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Created"})
		return
	}
	c.JSON(http.StatusOK, task)
}

//GetTaskDetailById godoc
//@Summary Show task detail including project, label, status based on Task Id
//@Tags tasks
//@ID get-string-by-int
//@Accept  json
//@Produce  json
//@Param task_id path int true "Task ID"
//@success 200 {object} models.Task
//@Router /task/{task_id} [get]
func GetTaskById(c *gin.Context) {
	task_id := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Preload(("Project")).Preload(("Label")).Preload(("Status")).Preload(("User")).Where("task_id", task_id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

//GetTasks godoc
//@Summary Show a list of Tasks
//@Tags tasks
//@Accept  json
//@Produce  json
//@success 200 {object} models.Task
//@Router /task [get]
func GetAllTask(c *gin.Context) {
	var task []models.Task
	if err := config.DB.Preload("Project").Preload("Label").Preload("Status").Preload(("User")).Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

//UpdateTaskById godoc
//@Summary Update task or items by TaskId
//@Description Update by JSON Task
//@Tags task
//@Accept  json
//@Produce  json
//@Param task_id path int true "Task Id"
//@Param task body models.Task true "Update Task"
//@success 200 {object} models.Task
//@Router /task_id/{task_id} [put]
func UpdateTask(c *gin.Context) {
	task_id := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Preload(("Project")).Preload(("Label")).Preload(("Status")).Preload(("User")).Where("task_id", task_id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found"})
		return
	}
	c.ShouldBindJSON(&task)
	if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Updated!"})
		return
	}
	c.JSON(http.StatusOK, task)
}

//DeleteTaskById godoc
//@Summary Delete task including project, label, status by task id
//@Tags task
//@ID get-string-by-int
//@Accept  json
//@Produce  json
//@Param task_id path int true "Task ID"
//@success 200 {object} models.Task
//@Router /task/{task_id} [delete]
func DeleteTask(c *gin.Context) {
	task_id := c.Params.ByName("task_id")
	var task models.Task
	if err := config.DB.Where("task_id=?", task_id).Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Cannot Be Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task_id" + task_id: "is deleted"})
}
