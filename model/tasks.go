package model

type Task struct {
	ID    int     `json:"id_task"`
	Name  string  `json:"name"`
	Hours float64 `json:"hours"`
}
