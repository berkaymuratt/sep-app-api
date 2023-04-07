package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientInfo PatientInfo        `bson:"patient_info" json:"patient_info"`
	DoctorInfo
}
