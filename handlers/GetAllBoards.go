package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all boards called")
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	fmt.Printf("%+v\n", result)

	// var board types.Board
	// json.Unmarshal(body, &board)
	// fmt.Printf("%+v\n", board)
}
