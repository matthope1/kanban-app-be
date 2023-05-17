package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	fmt.Println("db init called....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	name := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	dbURL := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + name + " port=" + port + " sslmode=disable"
	fmt.Println("dbURL: ", dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
