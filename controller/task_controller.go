package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salsapunk/Gerenciador-de-Tarefas/model"
)

type taskController struct {
	// Usecase
}

func NewTaskController() taskController {
	return taskController{}
}

func (t *taskController) GetTasks(ctx *gin.Context) {
	tasks := []model.Task{
		{
			ID:    1,
			Name:  "Jogar",
			Hours: 1,
		},
	}

	ctx.JSON(http.StatusOK, tasks)
}
