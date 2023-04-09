package dtos

import (
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoctorDto struct {
	ID           primitive.ObjectID `json:"id"`
	UserId       string             `json:"user_id"`
	UserPassword string             `json:"user_password,omitempty"`
	DoctorInfo   models.DoctorInfo  `json:"doctor_info"`
	Patients     []*PatientDto      `json:"patients,omitempty"`
}

func DoctorDtoFromDoctorDbResponse(doctorData *dbdtos.GetDoctorDbResponse) (*DoctorDto, error) {
	var patientsData []*PatientDto

	for _, patient := range doctorData.Patients {
		patientDto := PatientDto{
			ID:          patient.ID,
			DoctorId:    patient.DoctorId,
			UserId:      patient.UserId,
			PatientInfo: patient.PatientInfo,
		}
		patientsData = append(patientsData, &patientDto)
	}

	doctor := DoctorDto{
		ID:         doctorData.ID,
		UserId:     doctorData.UserId,
		DoctorInfo: doctorData.DoctorInfo,
		Patients:   patientsData,
	}
	return &doctor, nil
}
