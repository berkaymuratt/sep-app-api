package Dtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPatientResponse struct {
	ID          primitive.ObjectID `json:"id"`
	DoctorId    primitive.ObjectID `json:"doctor_id"`
	UserId      string             `json:"user_id"`
	PatientInfo models.PatientInfo `json:"patient_info"`
	Doctor      models.Doctor      `json:"doctor_models"`
}

type AddPatientRequest struct {
	DoctorId     primitive.ObjectID `json:"doctor_id"`
	UserId       string             `json:"user_id"`
	UserPassword string             `json:"user_password"`
	PatientInfo  models.PatientInfo `bson:"patient_info" json:"patient_info"`
}
