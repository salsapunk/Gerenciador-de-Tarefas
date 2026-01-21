package model

type Task struct {
	ID          int     `json:"id_task"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Time        float64 `json:"hours"`
	Done        bool    `json:"done"`
	CreatedAt   string  `json:"created_at"`
}
