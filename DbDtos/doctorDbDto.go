package DbDtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetDoctorDbResponse struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserId     string             `bson:"user_id"`
	DoctorInfo models.DoctorInfo  `bson:"doctor_info"`
	Patients   []models.Patient   `bson:"patients"`
}
