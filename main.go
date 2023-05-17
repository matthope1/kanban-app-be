package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pong"})
}

func main() {
	fmt.Println("hello this is main speaking")
	router := gin.Default()
	router.GET("/ping", ping)

	router.Run("localhost:8080")
}
