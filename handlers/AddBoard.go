package handlers

import (
	"fmt"
	"io/ioutil"
	"kanban-app-be/auth0"
	"log"
	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
)

func (h handler) AddBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add board called")
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	userInfo := auth0.GetUserInfo(token)
	fmt.Println("user email:", userInfo.Email)
	// print data from request
	fmt.Println("request body:", r.Body)

	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// TODO:
	// 1. get user id & board info from req
	// 2. Create new board obj with info
	// 3. commit to db
	// 4. send success response
	// 5. err handling

	fmt.Println("body", body)
}
