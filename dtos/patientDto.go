package dtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientDto struct {
	ID           primitive.ObjectID `json:"id"`
	DoctorId     primitive.ObjectID `json:"doctor_id"`
	UserId       string             `json:"user_id"`
	UserPassword string             `json:"user_password,omitempty"`
	PatientInfo  models.PatientInfo `json:"patient_info"`
}
