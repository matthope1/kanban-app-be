package main

import (
	"fmt"
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello this is main speaking")
	DB := db.Init()
	h := handlers.New(DB)
	fmt.Println(h)

	router := mux.NewRouter()

	router.HandleFunc("/ping", h.Ping).Methods(http.MethodGet)
	router.HandleFunc("/allBoards", h.GetAllBoards).Methods(http.MethodGet)
	router.HandleFunc("/addBoard", h.AddBoard).Methods(http.MethodPost)
	router.HandleFunc("/addColumn", h.AddColumn).Methods(http.MethodPost)

	// router.GET("/allBoards")
	http.ListenAndServe("localhost:8080", router)

}
