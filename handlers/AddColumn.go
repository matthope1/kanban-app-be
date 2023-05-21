package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddColumn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update column called")
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	// TODO:
	// 1. get board id from req
	// 2. create new column obj & give it the board id
	// 3. commit to db
	// 4. send success response

	fmt.Println("body", body)
}
