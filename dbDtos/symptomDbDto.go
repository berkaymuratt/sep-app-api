package dbDtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetSymptomDbResponse struct {
	ID            primitive.ObjectID `bson:"_id"`
	BodyPartId    primitive.ObjectID `bson:"_body_part_id"`
	Name          string             `bson:"name"`
	PainIntensity int                `bson:"pain_intensity"`
	BodyParts     []models.BodyPart  `bson:"body_parts"`
}
