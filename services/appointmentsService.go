package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//go:generate mockgen -destination=../mocks/service/mockAppointmentsService.go -package=services github.com/berkaymuratt/sep-app-api/services AppointmentsServiceI
type AppointmentsServiceI interface {
	GetAppointmentById(appointmentId primitive.ObjectID) (*dtos.AppointmentDto, error)
	GetAppointments(doctorId *primitive.ObjectID, patientId *primitive.ObjectID) ([]*dtos.AppointmentDto, error)
	AddAppointment(newAppointment models.Appointment) error
	UpdateAppointmentDate(appointmentId primitive.ObjectID, newDate time.Time) error
	DeleteAppointment(appointmentId primitive.ObjectID) error
	IsDateAvailable(doctorId primitive.ObjectID, patientId primitive.ObjectID, date time.Time) bool
}

type AppointmentsService struct {
	AppointmentsServiceI
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

func (service AppointmentsService) GetAppointments(doctorId *primitive.ObjectID, patientId *primitive.ObjectID) ([]*dtos.AppointmentDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cursor *mongo.Cursor
	var err error

	if doctorId != nil {
		cursor, err = service.getAppointmentsCursor(ctx, "_doctor_id", doctorId)
	} else if patientId != nil {
		cursor, err = service.getAppointmentsCursor(ctx, "_patient_id", patientId)
	} else {
		return nil, errors.New("missing ids")
	}

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

func (service AppointmentsService) AddAppointment(newAppointment models.Appointment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := configs.GetCollection("appointments").InsertOne(ctx, newAppointment)
	return err
}

func (service AppointmentsService) UpdateAppointmentDate(appointmentId primitive.ObjectID, newDate time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"date": newDate,
		},
	}

	_, err := configs.GetCollection("appointments").UpdateByID(ctx, appointmentId, update)
	return err
}

func (service AppointmentsService) DeleteAppointment(appointmentId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := configs.GetCollection("appointments").DeleteOne(ctx, bson.M{"_id": appointmentId})
	return err
}

func (service AppointmentsService) IsDateAvailable(doctorId primitive.ObjectID, patientId primitive.ObjectID, date time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("appointments")

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"$or": bson.A{
					bson.M{
						"$and": bson.A{
							bson.M{
								"_doctor_id": doctorId,
							},
							bson.M{
								"date": bson.M{
									"$gt": date.Add(-30 * time.Minute),
									"$lt": date.Add(30 * time.Minute),
								},
							},
						},
					},
					bson.M{
						"$and": bson.A{
							bson.M{
								"_patient_id": patientId,
							},
							bson.M{
								"date": bson.M{
									"$gt": date.Add(-30 * time.Minute),
									"$lt": date.Add(30 * time.Minute),
								},
							},
						},
					},
				},
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return false
	}

	var appointments []*models.Appointment
	if err := cursor.All(ctx, &appointments); err != nil {
		return false
	}

	return len(appointments) == 0
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
		bson.M{
			"$match": bson.M{
				"date": bson.M{
					"$gt": time.Now(),
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"date": 1,
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
