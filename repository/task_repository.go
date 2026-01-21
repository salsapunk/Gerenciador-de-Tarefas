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
	query := "SELECT id, task_name, task_description, task_time, done, created_at FROM tasks"
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
			&taskObj.Description,
			&taskObj.Time,
			&taskObj.Done,
			&taskObj.CreatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	err = rows.Close()
	if err != nil {
		fmt.Println(err)
		return []model.Task{}, err
	}

	return taskList, nil
}

func (tr *TaskRepository) CreateTask(task model.Task) (int, error) {
	var id int

	query, err := tr.connection.Prepare("INSERT INTO tasks(task_name, task_description, task_time, done, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(task.Name, task.Description, task.Time, task.Done, task.CreatedAt).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (tr *TaskRepository) GetTaskById(id_task int) (*model.Task, error) {
	query, err := tr.connection.Prepare("SELECT * FROM tasks WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var task model.Task
	err = query.QueryRow(id_task).Scan(
		&task.ID,
		&task.Name,
		&task.Description,
		&task.Time,
		&task.Done,
		&task.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	err = query.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &task, nil
}

// func (tr *TaskRepository) GetTaskByName(task_name string) (*model.Task, error) {
//	query, err := tr.connection.Prepare("SELECT * FROM tasks WHERE id = $1")
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}

//	var task model.Task
//	err = query.QueryRow(task_name).Scan(
//		&task.ID,
//		&task.Name,
//		&task.Hours,
//	)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, nil
//		}

//		return nil, err
//	}

//	err = query.Close()
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}

//	return &task, nil
//}
