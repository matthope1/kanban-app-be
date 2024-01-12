package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"kanban-app-be/auth0"
	"kanban-app-be/types"

	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	coll := h.DB.Database("kanban").Collection("boards")
	fmt.Println("all boards... \n user email: ", userInfo.Email)
	filter := bson.D{{"user_email", userInfo.Email}}
	// Retrieves the first matching document
	var result types.Board
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	// Prints a message if no documents are matched or if any
	// other errors occur during the operation
	fmt.Println("result: ", result)
	if err != nil {
		panic(err)
	}

	boardsJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("boards json")
	fmt.Println(string(boardsJson))

	w.WriteHeader(http.StatusCreated)
	w.Write(boardsJson)
}
