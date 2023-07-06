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

func (h handler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
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

	// // get user data from auth0

	// // TODO:
	// // 1. get user email from req  (how can we use auth0 to ensure that the requests are coming from the correct user?)
	fmt.Println("user info email", userInfo.Email)

	// // 2. get all boards from database for the current user
	// testing pgx database connection

	// // 3. return boards in success response
	// // 4. handle errors and send appropriate response

	// testing gorm again
	// var result string
	// h.DB.Raw("SELECT user_email FROM board WHERE id = ?", 1).Scan(&result)
	// fmt.Println("my board:", result)

	// TODO: add error handling, ensure that user email exists

	// TODO: turn this database call into its own standalone function

	// var boards []types.Board

	// h.DB.Raw("SELECT * from board WHERE user_email= ?", userInfo.Email).Scan(&boards)

	// print data for all boards

	// for i, board := range boards {
	// 	fmt.Println("board #: ", i)
	// 	fmt.Println("my board id:", board.ID)
	// 	fmt.Println("my board title:", board.Title)
	// 	fmt.Println("my board email:", board.UserEmail)
	// 	fmt.Println("my board status:", board.Status)
	// }

	newBoards := db.GetBoardsByEmail(h.DB, userInfo.Email)

	// print data for all boards
	fmt.Println("testing getboardsbyemail function...")
	for i, board := range newBoards {
		fmt.Println("board #: ", i)
		fmt.Println("my board id:", board.ID)
		fmt.Println("my board title:", board.Title)
		fmt.Println("my board email:", board.UserEmail)
		fmt.Println("my board status:", board.Status)
	}

	// TODO: get all boards where user email is userInfo.Email
	// You can have functions for each of these db operations rather than putting all of the queries in this function
	// TODO: function for getting all user boards,
	// TODO: function for getting all columns for a board id
	// TODO: function for getting all

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("response from  get all boards")
}
