package main

import (
	"context"
	"encoding/json"
	"fmt"
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"kanban-app-be/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Restaurant struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	RestaurantId string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       []interface{}
}

func main() {
	fmt.Println("starting program")
	// testing mongo
	mongoDB := db.InitMongoDb()

	// begin findOne
	coll := mongoDB.Database("sample_restaurants").Collection("restaurants")
	filter := bson.D{{"name", "Bagels N Buns"}}

	var result Restaurant
	var err error
	err = coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	// end findOne

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)

	fmt.Println("ending program")

	return

	// end testing mongo

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
