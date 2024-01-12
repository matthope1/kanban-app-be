package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDb() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoDbPass := os.Getenv("MONGO_DB_PASS")
	fmt.Println("mongoDbPass", mongoDbPass)

	URI := "mongodb+srv://root:" + mongoDbPass + "@cluster0.1vieuux.mongodb.net/?retryWrites=true&w=majority"
	fmt.Println("URI: ", URI)
	opts := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println("error connecting to mongo")
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}
