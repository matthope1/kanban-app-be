package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"kanban-app-be/auth0"
	"kanban-app-be/types"
	"net/http"
	"os"

	middleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/jackc/pgx/v5"
)

func (h handler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	token, _ := middleware.AuthHeaderTokenExtractor(r)
	fmt.Println("token from get all boards request..", token)
	// TODO: import local auth0 package and get user info from auth0
	userInfo := auth0.GetUserInfo(token)

	// allow all origins for now
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	// // get user data from auth0

	// // TODO:
	// // 1. get user email from req  (how can we use auth0 to ensure that the requests are coming from the correct user?)
	fmt.Println("user info email", userInfo.Email)

	// // 2. get all boards from database for the current user
	// testing pgx database connection

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// TODO: lets build a nice db url

	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println("connecting to this db url", dbUrl)

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var user_email string
	err = conn.QueryRow(context.Background(), "select user_email from board where id=$1", 1).Scan(&user_email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	fmt.Println("user email from database", user_email)

	// // 3. return boards in success response
	// // 4. handle errors and send appropriate response

	// testing gorm again
	// var result string
	// h.DB.Raw("SELECT user_email FROM board WHERE id = ?", 1).Scan(&result)
	// fmt.Println("my board:", result)

	// TODO: add error handling, ensure that user email exists
	var myBoard types.Board
	h.DB.Raw("SELECT * from board WHERE user_email= ?", userInfo.Email).Scan(&myBoard)
	fmt.Println("my board id:", myBoard.ID)
	fmt.Println("my board title:", myBoard.Title)
	fmt.Println("my board email:", myBoard.UserEmail)
	fmt.Println("my board status:", myBoard.Status)

	// TODO: get all boards where user email is userInfo.Email
	// TODO: get all other related information, columns, tasks, subtasks...

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("response from  get all boards")
}
