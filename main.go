package main

import (
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"kanban-app-be/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)
	router.Use(middleware.LoggingMiddleware)

	// applying middleware to CRUD routes
	api := router.PathPrefix("").Subrouter()
	api.Use(middleware.EnsureValidToken)

	// only signed in users can call these functions
	api.HandleFunc("/allBoards", h.AllBoards).Methods(http.MethodGet)
	api.HandleFunc("/addBoard", h.AddBoard).Methods(http.MethodPost)
	api.HandleFunc("/updateBoard", h.UpdateBoard).Methods(http.MethodPost)

	api.HandleFunc("/addColumn", h.AddColumn).Methods(http.MethodPost)
	api.HandleFunc("/updateColumn", h.UpdateColumn).Methods(http.MethodPost)

	api.HandleFunc("/addTask", h.AddTask).Methods(http.MethodPost)
	api.HandleFunc("/updateTask", h.UpdateTask).Methods(http.MethodPost)

	api.HandleFunc("/addSubtask", h.AddSubtask).Methods(http.MethodPost)
	api.HandleFunc("/updateSubtask", h.UpdateSubtask).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		// Debug: true,
	})

	handler := c.Handler(router)

	http.ListenAndServe("localhost:8080", handler)
}
