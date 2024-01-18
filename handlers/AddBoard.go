package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kanban-app-be/types"
	"log"
	"net/http"
)

func (h handler) AddBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add board called")

	coll := h.DB.Database("kanban").Collection("boards")

	// userInfo := auth0.GetUserInfo(token)

	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var board types.Board
	err = json.Unmarshal(body, &board)

	result, err := coll.InsertOne(context.TODO(), board)
	fmt.Println("result", result)

	// var result map[string]interface{}

	// if err := json.Unmarshal(body, &result); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("result", result)

	// boardID := result["boardId"]
	// taskID := result["taskId"]

	// get board id and task id from request body

	if err != nil {
		log.Fatal(err)
	}

	// 4. send success response
	// 5. err handling
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Successfully added board")
}
