package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salsapunk/Gerenciador-de-Tarefas/controller"
)

func main() {
	server := gin.Default()

	TaskController := controller.NewTaskController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/tasks", TaskController.GetTasks)

	server.Run(":8080")
}
