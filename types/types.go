package types

import (
	"time"

	"gorm.io/gorm"
)

type UserInfo struct {
	Sub           string
	Nickname      string
	Name          string
	Picture       string
	UpdatedAt     string
	Email         string
	EmailVerified string
}

type Board struct {
	gorm.Model
	ID        int
	UserEmail string    `json:"user_email"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Columns   []Column  `json:"columns"`
}

type Column struct {
	gorm.Model
	ID        int
	Title     string    `json:"title"`
	BoardId   int       `json:"board_id"`
	CreatedAt time.Time `json:"created_at"`
	Tasks     []Task    `json:"tasks"`
}

type Task struct {
	gorm.Model
	ID        int
	Status    string    `json:"status"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	ColumnId  int       `json:"column_id"`
	BoardId   int       `json:"board_id"`
	CreatedAt time.Time `json:"created_at"`
	Subtasks  []Subtask `json:"subtasks"`
}

type Subtask struct {
	gorm.Model
	ID         int
	TaskId     int       `json:"task_id"`
	BoardId    int       `json:"board_id"`
	Desc       string    `json:"desc"`
	IsComplete bool      `json:"is_complete"`
	CreatedAt  time.Time `json:"created_at"`
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
