package dtos

import (
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReportDto struct {
	ID               primitive.ObjectID `json:"id"`
	Doctor           *DoctorDto         `json:"doctor,omitempty"`
	Patient          *PatientDto        `json:"patient,omitempty"`
	Symptoms         []*SymptomDto      `json:"symptoms,omitempty"`
	PossibleDiseases []*DiseaseDto      `json:"possible_diseases,omitempty"`
	DoctorFeedback   string             `json:"doctor_feedback"`
	PatientNote      string             `json:"patient_note"`
	CreatedAt        time.Time          `json:"created_at"`
}

func ReportDtoFromReportDbResponse(reportData *dbdtos.GetReportDbResponse, symptomsDtos []*SymptomDto, diseasesDtos []*DiseaseDto) (*ReportDto, error) {
	doctor := reportData.Doctors[0]
	patient := reportData.Patients[0]

	doctorDto := DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	patientDto := PatientDto{
		ID:          patient.ID,
		DoctorId:    patient.DoctorId,
		UserId:      patient.UserId,
		PatientInfo: patient.PatientInfo,
	}

	reportDto := ReportDto{
		ID:               reportData.ID,
		Doctor:           &doctorDto,
		Patient:          &patientDto,
		Symptoms:         symptomsDtos,
		PossibleDiseases: diseasesDtos,
		DoctorFeedback:   reportData.DoctorFeedback,
		PatientNote:      reportData.PatientNote,
		CreatedAt:        reportData.CreatedAt,
	}

	return &reportDto, nil
}
