package dbdtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPatientDbResponse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	DoctorId     primitive.ObjectID `bson:"_doctor_id"`
	UserId       string             `bson:"user_id"`
	UserPassword string             `bson:"user_password"`
	PatientInfo  models.PatientInfo `bson:"patient_info"`
	Doctors      []models.Doctor    `bson:"doctors"`
}
