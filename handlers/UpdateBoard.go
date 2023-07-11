package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"kanban-app-be/auth0"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

func (h handler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update board called")
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	userInfo := auth0.GetUserInfo(token)
	fmt.Println("user email:", userInfo.Email)

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("domp", string(requestDump))

	// TODO: look for the board id
	// only allow the update to go through if the user owns the board

	// allow all origins
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	fmt.Println("update board called")
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	// TODO:
	// 1. get board id & updated fields from req
	// 2. retreive board obj from db using id
	// 3. update fields & commit to db
	// 4. send response back based on success of previous steps

	fmt.Println("body", string(body))
	w.WriteHeader(http.StatusCreated)
	// respond with the board id
	// w.Write([]byte("this is the update board response"))
	json.NewEncoder(w).Encode("this is the update board response")
}
