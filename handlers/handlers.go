package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	DB *mongo.Client
}

func New(db *mongo.Client) handler {
	return handler{db}
}
