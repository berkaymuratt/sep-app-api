package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AppointmentDto struct {
	ID          primitive.ObjectID `json:"id"`
	ReportId    primitive.ObjectID `json:"report_id"`
	Doctor      *DoctorDto         `json:"doctor,omitempty"`
	Patient     *PatientDto        `json:"patient,omitempty"`
	Symptoms    []*SymptomDto      `json:"symptoms"`
	PatientNote string             `bson:"patient_note"`
	Date        time.Time          `json:"date"`
}
