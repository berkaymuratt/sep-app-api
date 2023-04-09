package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Symptom struct {
	ID         primitive.ObjectID `bson:"_id"`
	BodyPartId primitive.ObjectID `bson:"_body_part_id"`
	Name       string             `bson:"name"`
	Level      int                `bson:"level"`
}
