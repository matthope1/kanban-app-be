package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// middleware "github.com/auth0/go-jwt-middleware/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (h handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete task called")

	// token, _ := middleware.AuthHeaderTokenExtractor(r)
	// userInfo := auth0.GetUserInfo(token)

	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	// Now you can access the "boardId" field
	boardID := result["boardId"]
	taskID := result["taskId"]

	fmt.Println("boardId:", boardID)
	fmt.Println("taskId:", taskID)

	// get board id and task id from request body

	coll := h.DB.Database("kanban").Collection("boards")
	filter := bson.M{"_id": taskID}

	deleteCount, err := coll.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document Count: ", deleteCount)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Task updated successfully")

}
