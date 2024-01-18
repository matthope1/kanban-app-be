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

	// begin
	// begin find
	coll := h.DB.Database("kanban").Collection("boards")

	// Creates a query filter to match documents in which the "cuisine"
	// is "Italian"
	filter := bson.D{{"user_email", userInfo.Email}}

	// Retrieves documents that match the query filer
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// end find

	var results []types.Board
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Prints the results of the find operation as structs
	for _, result := range results {
		cursor.Decode(&result)
		// output, err := json.MarshalIndent(result, "", "    ")
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%s\n", output)
	}

	boardsJson, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("boards json")
	// fmt.Println(string(boardsJson))

	w.WriteHeader(http.StatusCreated)
	w.Write(boardsJson)
}
