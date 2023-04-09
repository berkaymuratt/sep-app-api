package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReportsService struct{}

func NewReportsService() ReportsService {
	return ReportsService{}
}

func (service ReportsService) GetReport(reportId primitive.ObjectID) (*dtos.ReportDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	coll := configs.GetCollection("reports")
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
				"from":         "symptoms",
				"localField":   "_symptom_ids",
				"foreignField": "_id",
				"as":           "symptoms",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "diseases",
				"localField":   "_possible_disease_ids",
				"foreignField": "_id",
				"as":           "possible_diseases",
			},
		},
		bson.M{
			"$match": bson.M{
				"_id": reportId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetReportDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("cannot found any report")
	}

	if len(result) != 1 && len(result[0].Doctors) != 1 {
		return nil, errors.New("doctor cannot found")
	}

	if len(result) != 1 && len(result[0].Patients) != 1 {
		return nil, errors.New("patient cannot found")
	}

	data := result[0]
	doctor := data.Doctors[0]
	patient := data.Patients[0]

	doctorDto := dtos.DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	patientDto := dtos.PatientDto{
		ID:          patient.ID,
		DoctorId:    patient.DoctorId,
		UserId:      patient.UserId,
		PatientInfo: patient.PatientInfo,
	}

	report := dtos.ReportDto{
		ID:      data.ID,
		Doctor:  &doctorDto,
		Patient: &patientDto,
		//Symptoms:         data.Symptoms,
		//PossibleDiseases: data.PossibleDiseases,
		DoctorFeedback: data.DoctorFeedback,
		PatientNote:    data.PatientNote,
		CreatedAt:      data.CreatedAt,
	}

	return &report, nil
}
