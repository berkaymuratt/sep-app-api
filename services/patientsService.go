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

type PatientsService struct{}

func NewPatientsService() PatientsService {
	return PatientsService{}
}

func (service PatientsService) GetPatients() ([]*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getPatientsCursor(ctx, "", "")

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var patients []*dtos.PatientDto

	for _, patientData := range result {
		patientDto, err := dtos.PatientDtoFromPatientDbResponse(patientData)

		if err != nil {
			return nil, err
		}

		patients = append(patients, patientDto)
	}

	return patients, err
}

func (service PatientsService) GetPatientsByDoctorId(doctorId primitive.ObjectID) ([]*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getPatientsCursor(ctx, "_doctor_id", doctorId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var patients []*dtos.PatientDto

	for _, patientData := range result {
		patientDto, err := dtos.PatientDtoFromPatientDbResponse(patientData)

		if err != nil {
			return nil, err
		}

		patients = append(patients, patientDto)
	}

	return patients, err
}

func (service PatientsService) GetPatientById(patientId primitive.ObjectID) (*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getPatientsCursor(ctx, "_id", patientId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 && len(result[0].Doctors) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	patientData := result[0]
	return dtos.PatientDtoFromPatientDbResponse(patientData)
}

func (service PatientsService) AddPatient(patient models.Patient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := hashPassword(patient.UserPassword)

	if err != nil {
		return err
	}

	patient.UserPassword = hashedPassword

	coll := configs.GetCollection("patients")
	res, err := coll.InsertOne(ctx, patient)

	if err != nil && res == nil {
		return err
	}

	return nil
}

func (service PatientsService) IsUserIdExist(userId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := bson.D{{"$match", bson.D{{"user_id", userId}}}}

	var err error
	var cursor *mongo.Cursor

	if cursor, err = configs.GetCollection("patients").Aggregate(ctx, pipeline); err != nil {
		return true
	}

	var result []dbdtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return true
	}

	return len(result) > 0
}

func (service PatientsService) getPatientsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
	coll := configs.GetCollection("patients")
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "doctors",
				"localField":   "_doctor_id",
				"foreignField": "_id",
				"as":           "doctors",
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
