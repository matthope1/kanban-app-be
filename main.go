package main

import (
	"context"
	"fmt"
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"kanban-app-be/middleware"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type Subtask struct {
	ID         primitive.ObjectID `bson:"_id"`
	Desc       string             `bson:"desc"`
	IsComplete bool               `bson:"is_complete"`
	CreatedAt  time.Time          `bson:"created_at"`
}

type Column struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Desc      string             `bson:"desc"`
	CreatedAt time.Time          `bson:"created_at"`
	Tasks     []Task             `bson:"tasks"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	Status    string             `bson:"status"`
	Title     string             `bson:"title"`
	Desc      string             `bson:"desc"`
	CreatedAt time.Time          `bson:"created_at"`
	Subtasks  []Subtask          `bson:"subtasks"`
}

type Board struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	UserEmail string             `bson:"user_email"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	Columns   []Column           `bson:"columns"`
}

func main() {
	fmt.Println("starting program")
	// testing mongo
	mongoDB := db.InitMongoDb()

	// Specify the database and collection
	database := mongoDB.Database("kanban")
	collection := database.Collection("boards")

	// Find all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor and print each document
	for cursor.Next(context.TODO()) {
		var result Board
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Found document: %+v\n", result)
		fmt.Println("title: ", result.Title)
		fmt.Println("user_email: ", result.UserEmail)
		fmt.Println("status: ", result.Status)
		fmt.Println("created_at: ", result.CreatedAt)
		fmt.Println("columns: ", result.Columns)
		fmt.Println("=====================================")

	}

	// Check for errors from iterating over the cursor
	if err := cursor.Err(); err != nil {
		panic(err)
	}

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
