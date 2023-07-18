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

// reuse code from AddColumn.go to create an AddTask handler
func (h handler) AddTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add task called")
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

	var task types.Task
	err = json.Unmarshal(body, &task)

	if err != nil {
		fmt.Println("error unmarshalling task for user: ", userInfo.Email)
		log.Fatalln(err)
		return
	}

	boardOwner := db.GetBoardOwnerById(h.DB, task.BoardId)
	if boardOwner != userInfo.Email {
		fmt.Println("user does not own the board that this task belongs to ")
		json.NewEncoder(w).Encode("User does not own the board that this task belongs to")
		return
	}

	db.AddTask(h.DB, task)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Task added successfully")

}
