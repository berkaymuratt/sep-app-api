package dbDtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetReportDbResponse struct {
	ID                 primitive.ObjectID   `bson:"_id"`
	DoctorId           primitive.ObjectID   `bson:"_doctor_id"`
	PatientId          primitive.ObjectID   `bson:"_patient_id"`
	SymptomIds         []primitive.ObjectID `bson:"_symptom_ids"`
	PossibleDiseaseIds []primitive.ObjectID `bson:"_possible_disease_ids"`
	DoctorFeedback     string               `bson:"doctor_feedback"`
	PatientNote        string               `bson:"patient_note"`
	CreatedAt          time.Time            `bson:"created_at"`

	Doctors          []models.Doctor  `bson:"doctors"`
	Patients         []models.Patient `bson:"patients"`
	Symptoms         []models.Symptom `bson:"symptoms"`
	PossibleDiseases []models.Disease `bson:"possible_diseases"`
}
