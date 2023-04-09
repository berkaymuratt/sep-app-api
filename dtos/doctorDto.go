package dtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoctorDto struct {
	ID           primitive.ObjectID `json:"id"`
	UserId       string             `json:"user_id"`
	UserPassword string             `json:"user_password,omitempty"`
	DoctorInfo   models.DoctorInfo  `json:"doctor_info"`
}
