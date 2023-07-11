package handlers

import (
	"encoding/json"
	"fmt"

	"kanban-app-be/auth0"
	"kanban-app-be/db"

	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

// TODO: rename to all data
func (h handler) AllBoards(w http.ResponseWriter, r *http.Request) {
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	userInfo := auth0.GetUserInfo(token)

	// allow all origins
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	// TODO: add error handling, ensure that user email exists

	// 2. get all boards from database for the current user
	boards := db.GetBoards(h.DB, userInfo.Email)

	// get all columns, tasks, and subtasks for each board
	for i, board := range boards {
		// get all columns for this board, and append to board struct
		columns := db.GetColumns(h.DB, board.ID)
		for j, column := range columns {
			boards[i].Columns = append(boards[i].Columns, column)

			fmt.Println("my column id:", column.ID)
			fmt.Println("my column title:", column.Title)
			fmt.Println("my column board id:", column.BoardId)
			fmt.Println("my column tasks :", column.Tasks)

			// for each column, get all tasks and append to column struct
			tasks := db.GetTasks(h.DB, column.ID)

			for k, task := range tasks {
				// append task to column struct
				board.Columns[j].Tasks = append(board.Columns[j].Tasks, task)

				// for each task, get all subtasks and append to task struct
				subtasks := db.GetSubTasks(h.DB, task.ID)

				for _, subtask := range subtasks {
					board.Columns[j].Tasks[k].Subtasks = append(board.Columns[j].Tasks[k].Subtasks, subtask)
				}
				// TODO: test adding the subtasks this way
				// board.Columns[j].Tasks[k].Subtasks = append(board.Columns[j].Tasks[k].Subtasks, subtasks...)
			}
		}
	}

	boardsJson, err := json.Marshal(boards)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("boards json")
	fmt.Println(string(boardsJson))

	w.WriteHeader(http.StatusCreated)
	w.Write(boardsJson)
}
