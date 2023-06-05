package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update column called")
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