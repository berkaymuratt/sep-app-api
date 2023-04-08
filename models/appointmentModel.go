package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Appointment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	DoctorId  primitive.ObjectID `bson:"_doctor_id" json:"doctorId"`
	PatientId primitive.ObjectID `bson:"_patient_id" json:"patientId"`
	ReportId  primitive.ObjectID `bson:"_report_id" json:"reportId"`
	Date      string             `bson:"date" json:"date"`
}
