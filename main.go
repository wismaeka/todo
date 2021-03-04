package main

import (
	"todo/config"
	"todo/controllers"
	_ "todo/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Todo API
// @version 1.0
// @description This is a basic API Todo of which has proprerties of Project, Label, Status, and Title using Gin and Gorm.

// @contact.name Todo.Support
// @contact.email wisma@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()
	config.ConnectDB()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/task", controllers.GetAllTask)
		v1.GET("/task/:task_id", controllers.GetTaskById)
		v1.POST("/task", controllers.CreateTask)
		v1.POST("/project", controllers.CreateProject)
		v1.POST("/label", controllers.CreateLabel)
		v1.PUT("/task/:task_id", controllers.UpdateTask)
		v1.DELETE("/task/:task_id", controllers.DeleteTask)
	}

	r.Run()

}
