package handlers

import (
	// "kanban-app-be/auth0"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	"net/http"
)

func (h handler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all boards called... look for access token in the headers...")

	// allow all origins for now 
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("response from all boards y'all...")


	// // Read to request body
	// defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var result map[string]any
	// json.Unmarshal([]byte(body), &result)
	// fmt.Printf("%+v\n", result)


	// // get access token from request

	// // get user data from auth0 

	// // TODO: 
	// // 1. get user id from req  (how can we use auth0 to ensure that the requests are coming from the correct user?)
	// // https://auth0.com/docs/quickstart/backend/golang/interactive check this link
	// // 2. get all boards from database for the current user
	// // 3. return boards in success response
	// // 4. handle errors and send appropriate response

	// // var board types.Board
	// // json.Unmarshal(body, &board)
	// // fmt.Printf("%+v\n", board)

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("response from  get all boards")
}
