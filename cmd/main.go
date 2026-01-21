package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salsapunk/Gerenciador-de-Tarefas/controller"
	"github.com/salsapunk/Gerenciador-de-Tarefas/db"
	"github.com/salsapunk/Gerenciador-de-Tarefas/repository"
	"github.com/salsapunk/Gerenciador-de-Tarefas/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	TaskRepository := repository.NewTaskRepository(dbConnection)
	// camada usecase
	TaskUseCase := usecase.NewTaskUseCase(TaskRepository)
	// camada controller
	TaskController := controller.NewTaskController(TaskUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/tasks", TaskController.GetTasks)
	server.POST("/task", TaskController.CreateTask)
	server.GET("task/:taskId", TaskController.GetTaskById)

	server.Run(":8080")
}
