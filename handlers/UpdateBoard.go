package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"kanban-app-be/auth0"
// 	"kanban-app-be/db"
// 	"kanban-app-be/types"

// 	middleware "github.com/auth0/go-jwt-middleware/v2"
// )

// func (h handler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("update board called")
// 	token, _ := middleware.AuthHeaderTokenExtractor(r)
// 	userInfo := auth0.GetUserInfo(token)

// 	// allow all origins
// 	w.Header().Add("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Credentials", "true")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 	// Read to request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// fmt.Println("body", string(body))

// 	var board types.Board
// 	err = json.Unmarshal(body, &board)
// 	if err != nil {
// 		fmt.Println("error unmarshalling board for user: ", userInfo.Email)
// 		log.Fatalln(err)
// 		return
// 	}

// 	fmt.Println("attempting to update board...")
// 	if userInfo.Email != board.UserEmail {
// 		fmt.Println("user email does not match board user email")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode("user email does not match board user email")
// 		return
// 	}

// 	db.UpdateBoard(h.DB, board)
// 	fmt.Println("after board update...")

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Successfully updated board")
// }
