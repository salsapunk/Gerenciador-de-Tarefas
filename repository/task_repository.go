package repository

import (
	"database/sql"
	"fmt"

	"github.com/salsapunk/Gerenciador-de-Tarefas/model"
)

type TaskRepository struct {
	connection *sql.DB
}

func NewTaskRepository(connection *sql.DB) TaskRepository {
	return TaskRepository{
		connection: connection,
	}
}

func (tr *TaskRepository) GetTask() ([]model.Task, error) {
	query := "SELECT id, task_name, hours FROM tasks"
	rows, err := tr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Task{}, err
	}

	var taskList []model.Task
	var taskObj model.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Name,
			&taskObj.Hours,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	rows.Close()
	return taskList, nil
}

func (tr *TaskRepository) CreateTask(task model.Task) (int, error) {
	var id int
	query, err := tr.connection.Prepare("INSERT INTO task(task_name, hours) VALUES($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(task.Name, task.Hours).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
