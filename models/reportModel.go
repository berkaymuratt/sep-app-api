package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Report struct {
	ID                 primitive.ObjectID   `bson:"_id"`
	DoctorId           primitive.ObjectID   `bson:"_doctor_id"`
	PatientId          primitive.ObjectID   `bson:"_patient_id"`
	SymptomIds         []primitive.ObjectID `bson:"_symptom_ids"`
	PossibleDiseaseIds []primitive.ObjectID `bson:"_possible_disease_ids"`
	DoctorFeedback     string               `bson:"doctor_feedback"`
	PatientNode        string               `bson:"patient_note"`
	CreatedAt          time.Time            `bson:"created_at"`
}
