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

// func (h handler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Update column called")

// 	token, _ := middleware.AuthHeaderTokenExtractor(r)
// 	userInfo := auth0.GetUserInfo(token)

// 	// Read to request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var column types.Column
// 	err = json.Unmarshal(body, &column)

// 	if err != nil {
// 		fmt.Println("error unmarshalling column for user: ", userInfo.Email)
// 		log.Fatalln(err)
// 		return
// 	}

// 	fmt.Println("attempting to update column", column)

// 	// TODO: ensure that the board is owned by the user
// 	boardOwner := db.GetBoardOwnerById(h.DB, column.BoardId)
// 	fmt.Println("board owner", boardOwner)
// 	fmt.Println("user email from req", userInfo.Email)
// 	// print all column data
// 	fmt.Println("column", column)
// 	if boardOwner != userInfo.Email {
// 		fmt.Println("user does not own board")
// 		json.NewEncoder(w).Encode("User does not own this board")
// 		return
// 	}

// 	db.UpdateColumn(h.DB, column)
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Column updated successfully")

// }
