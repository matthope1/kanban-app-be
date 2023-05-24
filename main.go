package main

import (
	"fmt"
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"kanban-app-be/middleware"
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
	router.HandleFunc("/allTasks", h.AddColumn).Methods(http.MethodPost)

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.EnsureValidToken)


	// router.HandleFunc("/test", middleware.EnsureValidToken()(
	// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Content-Type", "application/json")

	// 		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	// 		w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
	// 	}),
	// ))


	// router.GET("/allBoards")
	http.ListenAndServe("localhost:8080", router)

}
