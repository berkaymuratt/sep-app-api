package dbdtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetSymptomDbResponse struct {
	ID         primitive.ObjectID `bson:"_id"`
	BodyPartId primitive.ObjectID `bson:"_body_part_id"`
	Name       string             `bson:"name"`
	Level      int                `bson:"level"`
	BodyParts  []models.BodyPart  `bson:"body_parts"`
}
