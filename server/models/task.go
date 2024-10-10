package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	IsDone bool               `json:"isdone"`
}
