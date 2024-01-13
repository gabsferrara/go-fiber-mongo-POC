package tags

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tags struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}
