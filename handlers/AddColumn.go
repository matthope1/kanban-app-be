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
	// 1.
	// 2.
	// 3.
	// 4.

	fmt.Println("body", body)
}
