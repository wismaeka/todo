package main

import (
	"todo/config"
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/task", controllers.GetAllTask)
	r.GET("/task/:task_id", controllers.GetTaskById)
	r.POST("/task", controllers.CreateTask)
	r.POST("/project", controllers.CreateProject)
	r.POST("/label", controllers.CreateLabel)
	r.PUT("/task/:task_id", controllers.UpdateTask)
	r.DELETE("/task/:task_id", controllers.DeleteTask)
	config.ConnectDB()
	r.Run(":8081")

}
