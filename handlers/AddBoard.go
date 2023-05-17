package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add board called")
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("body", body)
}
