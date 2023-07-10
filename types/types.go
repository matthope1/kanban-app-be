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
	UserEmail string
	Title     string
	Status    string
	CreatedAt time.Time
	Columns   []Column
}

type Column struct {
	gorm.Model
	ID        int
	Title     string
	BoardId   int
	CreatedAt time.Time
	Tasks     []Task
}

type Task struct {
	gorm.Model
	ID        int
	Status    string
	Desc      string
	ColumnId  int
	CreatedAt time.Time
	Subtasks  []Subtask
}

type Subtask struct {
	gorm.Model
	ID        int
	TaskId    int
	CreatedAt time.Time
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
