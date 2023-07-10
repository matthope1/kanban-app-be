package handlers

import (
	"encoding/json"
	"fmt"
	"kanban-app-be/auth0"
	"kanban-app-be/db"

	// "kanban-app-be/types"
	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

// TODO: rename to all data

func (h handler) AllBoards(w http.ResponseWriter, r *http.Request) {
	// get all boards is maybe not what we wanna call it.
	// We can call it get user data, and this function will return all boards columns etc...
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	// TODO: import local auth0 package and get user info from auth0
	userInfo := auth0.GetUserInfo(token)

	// allow all origins for now
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	fmt.Println("user info email", userInfo.Email)

	// // 3. return boards in success response
	// // 4. handle errors and send appropriate response

	// TODO: add error handling, ensure that user email exists

	// 2. get all boards from database for the current user
	boards := db.GetBoards(h.DB, userInfo.Email)
	// https://stackoverflow.com/questions/8270816/converting-go-struct-to-json

	// print data for all boards
	fmt.Println("testing getboardsbyemail function...")
	for i, board := range boards {
		// fmt.Println("board #: ", i)
		// fmt.Println("my board id:", board.ID)
		// fmt.Println("my board title:", board.Title)
		// fmt.Println("my board email:", board.UserEmail)
		// fmt.Println("my board status:", board.Status)
		// print all board data
		fmt.Println("board #: ", i)
		fmt.Println("my board id:", board.ID)
		fmt.Println("my board title:", board.Title)
		fmt.Println("my board email:", board.UserEmail)
		fmt.Println("my board status:", board.Status)

		// get all columns for this board, and append to board struct
		columns := db.GetColumns(h.DB, board.ID)
		// if there's columns add them to the board struct
		fmt.Println("columns length", len(columns))

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
			}
		}

		// create a struct that has a list of structs

		boardsJson, err := json.Marshal(boards)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(boardsJson))

		// TODO: function for getting all columns for a board id

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(string(boardsJson))
	}
}
