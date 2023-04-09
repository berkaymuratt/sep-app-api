package dtos

import (
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AppointmentDto struct {
	ID          primitive.ObjectID `json:"id"`
	ReportId    primitive.ObjectID `json:"report_id"`
	Doctor      *DoctorDto         `json:"doctor,omitempty"`
	Patient     *PatientDto        `json:"patient,omitempty"`
	Symptoms    []*SymptomDto      `json:"symptoms"`
	PatientNote string             `json:"patient_note"`
	Date        time.Time          `json:"date"`
}

func AppointmentDtoFromAppointmentResponse(appointmentData *dbdtos.GetAppointmentDbResponse, symptomsDtos []*SymptomDto) (*AppointmentDto, error) {
	doctorData := appointmentData.Doctors[0]
	doctorDto := DoctorDto{
		ID:         doctorData.ID,
		UserId:     doctorData.UserId,
		DoctorInfo: doctorData.DoctorInfo,
	}

	patientData := appointmentData.Patients[0]
	patientDto := PatientDto{
		ID:          patientData.ID,
		DoctorId:    patientData.DoctorId,
		UserId:      patientData.UserId,
		PatientInfo: patientData.PatientInfo,
	}

	appointmentDto := AppointmentDto{
		ID:          appointmentData.ID,
		ReportId:    appointmentData.ReportId,
		Doctor:      &doctorDto,
		Patient:     &patientDto,
		PatientNote: appointmentData.PatientNote,
		Symptoms:    symptomsDtos,
		Date:        appointmentData.Date,
	}

	return &appointmentDto, nil
}
