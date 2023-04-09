package dbdtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetDiseaseDbResponse struct {
	ID         primitive.ObjectID   `bson:"_id"`
	SymptomIds []primitive.ObjectID `bson:"_symptom_ids"`
	Name       string               `bson:"name"`
	Details    string               `bson:"details"`
}
