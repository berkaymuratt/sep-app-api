package models

import (
	models "github.com/berkaymuratt/sep-app-api/models/doctor"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPatientDbResponse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	DoctorId     primitive.ObjectID `bson:"_doctor_id"`
	UserId       string             `bson:"user_id"`
	UserPassword string             `bson:"user_password"`
	PatientInfo  PatientInfo        `bson:"patient_info"`
	Doctors      []models.Doctor    `bson:"doctors"`
}
