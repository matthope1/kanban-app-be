package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	Sub           string
	Nickname      string
	Name          string
	Picture       string
	UpdatedAt     string
	Email         string
	EmailVerified string
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
