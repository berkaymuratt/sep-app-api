package dtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SymptomDto struct {
	ID       primitive.ObjectID `json:"id"`
	BodyPart *models.BodyPart   `json:"body_part,omitempty"`
	Name     string             `json:"name"`
	Level    int                `json:"level"`
}
