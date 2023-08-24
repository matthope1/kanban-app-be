package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"kanban-app-be/auth0"
	"kanban-app-be/db"
	"kanban-app-be/types"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

func (h handler) AddColumn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add column called")

	token, _ := middleware.AuthHeaderTokenExtractor(r)
	userInfo := auth0.GetUserInfo(token)
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// print data from request
	fmt.Println("body", string(body))

	var column types.Column
	err = json.Unmarshal(body, &column)

	if err != nil {
		fmt.Println("error unmarshalling column for user: ", userInfo.Email)
		log.Fatalln(err)
		return
	}

	// TODO: ensure that the board is owned by the user
	boardOwner := db.GetBoardOwnerById(h.DB, column.BoardId)
	fmt.Println("board owner", boardOwner)
	fmt.Println("user email from req", userInfo.Email)
	// print all column data
	fmt.Println("column", column)
	if boardOwner != userInfo.Email {
		fmt.Println("user does not own board")
		json.NewEncoder(w).Encode("User does not own this board")
		return
	}

	db.AddColumn(h.DB, column, column.BoardId)

	// TODO:
	// 1. get board id from req
	// 2. create new column obj & add board id from step 1
	// 3. commit to db
	// 4. send success response
	// 5. handle errors and send appropriate response

	json.NewEncoder(w).Encode("Successfully added new column")
}
