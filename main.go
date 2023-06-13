package main

import (
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"kanban-app-be/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/ping", h.Ping).Methods(http.MethodGet)
	router.Use(middleware.LoggingMiddleware)

	// applying middleware to specific routes 
	api := router.PathPrefix("").Subrouter()
	api.Use(middleware.EnsureValidToken)

	// only signed in users can call these functions
	api.HandleFunc("/allBoards", h.GetAllBoards).Methods(http.MethodGet)
	api.HandleFunc("/addBoard", h.AddBoard).Methods(http.MethodPost)
	api.HandleFunc("/addColumn", h.AddColumn).Methods(http.MethodPost)

	api.HandleFunc("/updateBoard", h.GetAllBoards).Methods(http.MethodGet)
	api.HandleFunc("/updateColumn", h.GetAllBoards).Methods(http.MethodGet)

	api.HandleFunc("/updateUser", h.GetAllBoards).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8080", router)
}
