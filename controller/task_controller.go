package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salsapunk/Gerenciador-de-Tarefas/model"
	"github.com/salsapunk/Gerenciador-de-Tarefas/usecase"
)

type taskController struct {
	taskUseCase usecase.TaskUsecase
}

func NewTaskController(usecase usecase.TaskUsecase) taskController {
	return taskController{
		taskUseCase: usecase,
	}
}

func (t *taskController) GetTasks(ctx *gin.Context) {
	tasks, err := t.taskUseCase.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *taskController) CreateTask(ctx *gin.Context) {
	var task model.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err) // 400
		return
	}

	insertedTask, err := t.taskUseCase.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTask)
}
