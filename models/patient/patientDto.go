package models

import (
	models "github.com/berkaymuratt/sep-app-api/models/doctor"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPatientResponse struct {
	ID          primitive.ObjectID `json:"id"`
	DoctorId    primitive.ObjectID `json:"doctor_id"`
	UserId      string             `json:"user_id"`
	PatientInfo PatientInfo        `json:"patient_info"`
	Doctor      models.Doctor      `json:"doctor"`
}

type AddPatientRequest struct {
	DoctorId     primitive.ObjectID `json:"doctor_id"`
	UserId       string             `json:"user_id"`
	UserPassword string             `json:"user_password"`
	PatientInfo  PatientInfo        `bson:"patient_info" json:"patient_info"`
}
