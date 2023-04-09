package dbDtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetAppointmentDbResponse struct {
	ID          primitive.ObjectID   `bson:"_id"`
	DoctorId    primitive.ObjectID   `bson:"_doctor_id"`
	PatientId   primitive.ObjectID   `bson:"_patient_id"`
	ReportId    primitive.ObjectID   `bson:"_report_id"`
	SymptomIds  []primitive.ObjectID `bson:"_symptom_ids"`
	PatientNote string               `bson:"patient_note"`
	Date        time.Time            `bson:"date"`

	Doctors  []models.Doctor  `bson:"doctors"`
	Patients []models.Patient `bson:"patients"`
}
