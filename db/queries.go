package db

import (
	"kanban-app-be/types"

	"gorm.io/gorm"
)

func GetBoards(db *gorm.DB, userEmail string) []types.Board {
	var boards []types.Board
	db.Raw("SELECT * from board WHERE user_email = ?", userEmail).Scan(&boards)
	return boards
}

func GetColumns(db *gorm.DB, boardId int) []types.Column {
	var columns []types.Column
	db.Raw(`SELECT * from "column" c WHERE board_id = ?`, boardId).Scan(&columns)
	return columns
}

func GetTasks(db *gorm.DB, columnId int) []types.Task {
	var tasks []types.Task
	db.Raw("SELECT * from task WHERE column_id = ?", columnId).Scan(&tasks)
	return tasks
}

func GetSubTasks(db *gorm.DB, taskId int) []types.Subtask {
	var subtasks []types.Subtask
	db.Raw("SELECT * from subtask WHERE task_id = ?", taskId).Scan(&subtasks)
	return subtasks
}

// TODO: update functions
func UpdateBoard(db *gorm.DB, board types.Board) {
	db.Save(&board)
}
