package main

import (
	"fmt"
	"kanban-app-be/db"
	"kanban-app-be/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pong"})
}

func main() {
	fmt.Println("hello this is main speaking")
	DB := db.Init()
	handlers := handlers.New(DB)
	fmt.Println(handlers)
	router := gin.Default()
	router.GET("/ping", ping)

	router.Run("localhost:8080")
}
