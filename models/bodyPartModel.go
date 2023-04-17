package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BodyPart struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
}
