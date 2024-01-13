package tasks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Assign      primitive.ObjectID `json:"assign,omitempty" bson:"assign,omitempty"`
	Done        bool               `json:"done,omitempty" bson:"done,omitempty"`
}
