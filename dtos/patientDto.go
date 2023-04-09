package dtos

import (
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientDto struct {
	ID           primitive.ObjectID `json:"id"`
	DoctorId     primitive.ObjectID `json:"doctor_id"`
	UserId       string             `json:"user_id"`
	UserPassword string             `json:"user_password,omitempty"`
	PatientInfo  models.PatientInfo `json:"patient_info"`
	Doctor       *DoctorDto         `json:"doctor,omitempty"`
}

func PatientDtoFromPatientDbResponse(patientData *dbdtos.GetPatientDbResponse) (*PatientDto, error) {
	doctor := patientData.Doctors[0]
	doctorDto := DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	patientDTO := PatientDto{
		ID:          patientData.ID,
		DoctorId:    patientData.DoctorId,
		UserId:      patientData.UserId,
		PatientInfo: patientData.PatientInfo,
		Doctor:      &doctorDto,
	}

	return &patientDTO, nil
}
