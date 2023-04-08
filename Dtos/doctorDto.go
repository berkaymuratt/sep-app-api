package Dtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetDoctorResponse struct {
	ID         primitive.ObjectID `json:"id"`
	UserId     string             `json:"user_id"`
	DoctorInfo models.DoctorInfo  `json:"doctor_info"`
	Patients   []models.Patient   `json:"patients"`
}

type AddDoctorRequest struct {
	UserId       string            `json:"user_id"`
	UserPassword string            `json:"user_password"`
	DoctorInfo   models.DoctorInfo `json:"doctor_info"`
}
