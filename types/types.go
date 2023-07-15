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
	UserEmail string `json:"user_email"`
	Title     string
	Status    string `json:"status"`
	CreatedAt time.Time
	Columns   []Column
}

type Column struct {
	gorm.Model
	ID        int
	Title     string `json:"title"`
	BoardId   int    `json:"board_id"`
	CreatedAt time.Time
	Tasks     []Task
}

type Task struct {
	gorm.Model
	ID        int
	Status    string
	Desc      string
	ColumnId  int `json:"column_id"`
	CreatedAt time.Time
	Subtasks  []Subtask
}

type Subtask struct {
	gorm.Model
	ID         int
	TaskId     int `json:"task_id"`
	Desc       string
	IsComplete bool `json:"is_complete"`
	CreatedAt  time.Time
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
