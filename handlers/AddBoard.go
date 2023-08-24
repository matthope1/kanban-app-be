package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kanban-app-be/auth0"
	"kanban-app-be/db"
	"kanban-app-be/types"
	"log"
	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

func (h handler) AddBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add board called")
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	userInfo := auth0.GetUserInfo(token)

	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading request body")
		log.Fatalln(err)
		return
	}

	// print data from request
	fmt.Println("body", string(body))

	// TODO:
	// 1. get user email & board info from req
	// 2. Create new board obj with info
	var board types.Board
	err = json.Unmarshal(body, &board)

	if err != nil {
		fmt.Println("error unmarshalling board for user: ", userInfo.Email)
		log.Fatalln(err)
		return
	}

	fmt.Println("does it come with columns???? ---- len(columns): ", len(board.Columns))

	// 3. commit to db
	ID, addBoardErr := db.AddBoard(h.DB, board, userInfo.Email)
	// TODO: update the add board db call so that it returns the id of the board just created
	fmt.Println("just added a board with id:", ID)
	if addBoardErr != nil {
		fmt.Println("error adding board to db:", addBoardErr)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error adding board")
	}

	// TODO: if there's columns we also need to add them to the db
	// call db.AddColumn for each column :D

	// Adding all columns

	for _, column := range board.Columns {
		fmt.Println("adding column:", column.Title, "to board with id:", ID)
		// TODO: add error handling
		db.AddColumn(h.DB, column, ID)
	}

	// 4. send success response
	// 5. err handling
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Successfully added board")
}
