package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbDtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AppointmentsService struct {
	symptomsService SymptomsService
}

func NewAppointmentsService(symptomsService SymptomsService) AppointmentsService {
	return AppointmentsService{
		symptomsService: symptomsService,
	}
}

func (service AppointmentsService) GetAppointment(appointmentId primitive.ObjectID) (*dtos.AppointmentDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("appointments")

	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "doctors",
				"localField":   "_doctor_id",
				"foreignField": "_id",
				"as":           "doctors",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "patients",
				"localField":   "_patient_id",
				"foreignField": "_id",
				"as":           "patients",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "reports",
				"localField":   "_report_id",
				"foreignField": "_id",
				"as":           "reports",
			},
		},
		bson.M{
			"$match": bson.M{
				"_id": appointmentId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbDtos.GetAppointmentDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("appointment cannot found")
	}

	appointmentData := result[0]
	var symptomsData []dbDtos.GetSymptomDbResponse

	if symptomsData, err = service.symptomsService.GetSymptomsByIds(appointmentData.SymptomIds); err != nil {
		return nil, err
	}

	doctorData := appointmentData.Doctors[0]
	doctorDto := dtos.DoctorDto{
		ID:         doctorData.ID,
		UserId:     doctorData.UserId,
		DoctorInfo: doctorData.DoctorInfo,
	}

	patientData := appointmentData.Patients[0]
	patientDto := dtos.PatientDto{
		ID:          patientData.ID,
		DoctorId:    patientData.DoctorId,
		UserId:      patientData.UserId,
		PatientInfo: patientData.PatientInfo,
	}

	var symptomsDtos []*dtos.SymptomDto

	for _, symptom := range symptomsData {
		symptomDto := dtos.SymptomDto{
			ID:            symptom.ID,
			BodyPart:      &symptom.BodyParts[0],
			Name:          symptom.Name,
			PainIntensity: symptom.PainIntensity,
		}
		symptomsDtos = append(symptomsDtos, &symptomDto)
	}

	appointmentDto := dtos.AppointmentDto{
		ID:          appointmentData.ID,
		ReportId:    appointmentData.ReportId,
		Doctor:      &doctorDto,
		Patient:     &patientDto,
		PatientNote: appointmentData.PatientNote,
		Symptoms:    symptomsDtos,
		Date:        appointmentData.Date,
	}

	return &appointmentDto, err
}
