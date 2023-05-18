package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("body", body)
}
