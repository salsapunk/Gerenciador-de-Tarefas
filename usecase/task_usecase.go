package usecase

import (
	"github.com/salsapunk/Gerenciador-de-Tarefas/model"
	"github.com/salsapunk/Gerenciador-de-Tarefas/repository"
)

type TaskUsecase struct {
	repository repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) TaskUsecase {
	return TaskUsecase{
		repository: repo,
	}
}

func (tu *TaskUsecase) GetTasks() ([]model.Task, error) {
	return tu.repository.GetTask()
}

func (tu *TaskUsecase) CreateTask(task model.Task) (model.Task, error) {
	taskId, err := tu.repository.CreateTask(task)
	if err != nil {
		return model.Task{}, err
	}

	task.ID = taskId
	return task, nil
}
