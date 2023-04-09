package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (service AppointmentsService) GetAppointmentById(appointmentId primitive.ObjectID) (*dtos.AppointmentDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getAppointmentsCursor(ctx, "_id", appointmentId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetAppointmentDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("appointment cannot found")
	}

	appointmentData := result[0]

	var symptomsDto []*dtos.SymptomDto
	if symptomsDto, err = service.symptomsService.GetSymptomsByIds(appointmentData.SymptomIds); err != nil {
		return nil, err
	}

	return dtos.AppointmentDtoFromAppointmentResponse(appointmentData, symptomsDto)
}

func (service AppointmentsService) GetAppointmentByDoctor(doctorId primitive.ObjectID) ([]*dtos.AppointmentDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getAppointmentsCursor(ctx, "_doctor_id", doctorId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetAppointmentDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var appointmentsDtos []*dtos.AppointmentDto

	for _, appointmentData := range result {
		var symptomsDto []*dtos.SymptomDto
		if symptomsDto, err = service.symptomsService.GetSymptomsByIds(appointmentData.SymptomIds); err != nil {
			return nil, err
		}

		appointmentDto, err := dtos.AppointmentDtoFromAppointmentResponse(appointmentData, symptomsDto)

		if err != nil {
			return nil, err
		}

		appointmentsDtos = append(appointmentsDtos, appointmentDto)
	}

	return appointmentsDtos, err
}

func (service AppointmentsService) getAppointmentsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
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
	}

	if matchField != "" {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				matchField: matchValue,
			},
		})
	}

	return coll.Aggregate(ctx, pipeline)
}
