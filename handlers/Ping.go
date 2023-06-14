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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("pong")
}
