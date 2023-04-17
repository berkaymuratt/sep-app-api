package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Disease struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	SymptomIds []primitive.ObjectID `bson:"_symptom_ids"`
	Name       string               `bson:"name"`
	Details    string               `bson:"details"`
}
