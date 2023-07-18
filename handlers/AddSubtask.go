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

func (h handler) AddSubtask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add subtask called")
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
	// 2. Create subtask obj with info
	var subtask types.Subtask
	err = json.Unmarshal(body, &subtask)

	if err != nil {
		fmt.Println("error unmarshalling board for user: ", userInfo.Email)
		log.Fatalln(err)
		return
	}

	boardOwner := db.GetBoardOwnerById(h.DB, subtask.BoardId)
	if boardOwner != userInfo.Email {
		fmt.Println("user does not own the board that this subtask belongs to ")
		json.NewEncoder(w).Encode("User does not own the board that this task belongs to")
		return
	}

	// 3. commit to db
	db.AddSubtask(h.DB, subtask)

	// 4. send success response
	// 5. err handling
	json.NewEncoder(w).Encode("Successfully added subtask")
}
