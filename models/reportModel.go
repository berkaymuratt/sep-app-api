package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportModel struct {
	ID                 primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	PatientID          primitive.ObjectID   `bson:"_patient_id" json:"patient_id"`
	DoctorID           primitive.ObjectID   `bson:"_doctor_id" json:"doctor_id"`
	SymptomIDs         []primitive.ObjectID `bson:"symptom_ids" json:"symptom_ids"`
	PossibleDiseaseIDs []primitive.ObjectID `bson:"possible_disease_ids" json:"possible_disease_ids"`
}
