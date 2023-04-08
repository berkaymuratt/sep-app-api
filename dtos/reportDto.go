package dtos

import (
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReportDto struct {
	ID               primitive.ObjectID `json:"id"`
	Doctor           *DoctorDto         `json:"doctor,omitempty"`
	Patient          *PatientDto        `json:"patient,omitempty"`
	Symptoms         []*models.Symptom  `json:"symptoms,omitempty"`
	PossibleDiseases []*models.Disease  `json:"possible_diseases,omitempty"`
	DoctorFeedback   string             `json:"doctor_feedback"`
	PatientNote      string             `json:"patient_note"`
	CreatedAt        time.Time          `json:"created_at"`
}
