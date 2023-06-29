package handlers

import (
	"encoding/json"
	"fmt"
	"kanban-app-be/auth0"
	"kanban-app-be/types"

	// "io/ioutil"
	// "log"
	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

func (h handler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	fmt.Println("token from get all boards request..", token)
	// TODO: import local auth0 package and get user info from auth0
	userInfo := auth0.GetUserInfo(token)

	// allow all origins for now
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("response from all boards y'all...")

	// // Read to request body
	// defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var result map[string]any
	// json.Unmarshal([]byte(body), &result)
	// fmt.Printf("%+v\n", result)

	// // get user data from auth0

	// // TODO:
	// // 1. get user email from req  (how can we use auth0 to ensure that the requests are coming from the correct user?)
	fmt.Println("user info email", userInfo.Email)

	fmt.Println("db obj", h.DB)
	// // 2. get all boards from database for the current user

	// I don't think we need to do automigrate

	// h.DB.AutoMigrate(&types.Board{})

	// test creating a board

	// h.DB.Create(&types.Board{ID: 0, Email: "matt-hope@hotmail.com", Title: "Test Board", User_id: 0, Status: "Todo", Created_at: time.Now()})

	// test read
	var board types.Boards

	// find the first board with email = matt-hope@hotmail.com

	h.DB.First(board, "id = ?", 0)

	fmt.Println(board.Title)

	// // 3. return boards in success response
	// // 4. handle errors and send appropriate response

	// // var board types.Board
	// // json.Unmarshal(body, &board)
	// // fmt.Printf("%+v\n", board)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("response from  get all boards")
}
