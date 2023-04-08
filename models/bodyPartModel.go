package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BodyPart struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
