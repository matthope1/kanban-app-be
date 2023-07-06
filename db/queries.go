package db

import (
	"kanban-app-be/types"

	"gorm.io/gorm"
)

func GetBoardsByEmail(db *gorm.DB, userEmail string) []types.Board {
	var boards []types.Board
	db.Raw("SELECT * from board WHERE user_email= ?", userEmail).Scan(&boards)
	return boards
}
