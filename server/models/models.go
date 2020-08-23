package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task"`
	Status bool               `json:"status"`
    Date time.Time            `json:"date"`
}
