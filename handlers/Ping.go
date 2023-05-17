package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PONG")
	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("pong")
}
