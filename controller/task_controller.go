package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}

	var task model.Task
	if err := json.Unmarshal(body, &task); err != nil {
		fmt.Printf("unmarshal error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedTask, err := t.taskUseCase.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, insertedTask)
}

func (t *taskController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("taskId")
	if id == "" {
		response := model.Response{
			Message: "Id da tarefa não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id da tarefa tem que ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return

	}

	task, err := t.taskUseCase.GetTaskById(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if task == nil {
		response := model.Response{
			Message: "Id da tarefa não existe",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// func (t *taskController) GetTaskByName(ctx *gin.Context) {
//	taskName := ctx.Param("taskName")
//	if taskName == "" {
//		response := model.Response{
//			Message: "O nome da tarefa não pode estar vazio",
//		}
//		ctx.JSON(http.StatusBadRequest, response)
//		return
//	}

//	task, err := t.taskUseCase.GetTaskByName(taskName)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, err)
//		return
//	}

//	if task == nil {
//		response := model.Response{
//			Message: "O nome da tarefa não existe",
//		}
//		ctx.JSON(http.StatusNotFound, response)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, task)
//}
