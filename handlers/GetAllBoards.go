package handlers

import (
	"encoding/json"
	"fmt"
	"kanban-app-be/auth0"
	"kanban-app-be/types"

	// "io/ioutil"
	// "log"
	"net/http"

	middleware "github.com/auth0/go-jwt-middleware/v2"
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
	// testing boards struct

	// var board types.Board

	// h.DB.Unscoped().First(&board, 0)

	// h.DB.Raw("SELECT * FROM board where id = ?", 0).Scan(&board) // find board with integer primary key
	// fmt.Println("board user email", board.UserEmail)

	// Migrate the schema
	// h.DB.AutoMigrate(&types.Board{})
	// h.DB.AutoMigrate(&types.Column{})
	// h.DB.AutoMigrate(&types.Task{})
	// h.DB.AutoMigrate(&types.Subtask{})

	// var tables []string
	// if err := h.DB.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(tables); i++ {
	// 	fmt.Println("tables: ", tables[i])
	// }

	// // basic example product from gorm docs
	// // Migrate the schema
	h.DB.AutoMigrate(&types.Product{})

	// // Create
	// h.DB.Create(&types.Product{Code: "D42", Price: 100})

	// Read
	var product types.Product
	h.DB.First(&product, 0) // find product with integer primary key
	// h.DB.First(&product, "code = ?", "") // find product with code D42
	fmt.Println("product code: ", product.Code)
	// end testing

	// // 3. return boards in success response
	// // 4. handle errors and send appropriate response

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("response from  get all boards")
}
