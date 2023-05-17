package types

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
}

type Board struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	User_id    int       `json:"user_id"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"created_at"`
}

type Column struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Board_id   int       `json:"board_id"`
	Created_at time.Time `json:"created_at"`
}

type Task struct {
	ID         int       `json:"id"`
	Status     string    `json:"status"`
	Desc       string    `json:"desc"`
	Column_id  int       `json:"board_id"`
	Created_at time.Time `json:"created_at"`
}

type Subtask struct {
	ID         int       `json:"id"`
	Task_id    int       `json:"task_id"`
	Created_at time.Time `json:"created_at"`
}
