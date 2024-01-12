package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"kanban-app-be/auth0"
// 	"kanban-app-be/db"
// 	"kanban-app-be/types"
// 	"log"
// 	"net/http"

// 	middleware "github.com/auth0/go-jwt-middleware/v2"
// )

// func (h handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Update task called")

// 	token, _ := middleware.AuthHeaderTokenExtractor(r)
// 	userInfo := auth0.GetUserInfo(token)

// 	// Read to request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var task types.Task
// 	err = json.Unmarshal(body, &task)

// 	if err != nil {
// 		fmt.Println("error unmarshalling task for user: ", userInfo.Email)
// 		log.Fatalln(err)
// 		return
// 	}

// 	fmt.Println("attempting to update task", task)
// 	// ensure that the task belongs to a board owned by the user who made the request
// 	boardOwner := db.GetBoardOwnerById(h.DB, task.BoardId)
// 	if boardOwner != userInfo.Email {
// 		fmt.Println("user does not own board")
// 		json.NewEncoder(w).Encode("User does not own this board")
// 		return
// 	}

// 	db.UpdateTask(h.DB, task)
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Task updated successfully")

// }
