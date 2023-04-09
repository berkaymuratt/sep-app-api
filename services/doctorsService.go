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

type DoctorsService struct{}

func NewDoctorsService() DoctorsService {
	return DoctorsService{}
}

func (service DoctorsService) GetDoctors() ([]*dtos.DoctorDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getDoctorsCursor(ctx, "", "")

	if err != nil {
		return nil, err
	}

	var doctorsData []*dbdtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &doctorsData); err != nil {
		return nil, err
	}

	var doctors []*dtos.DoctorDto

	for _, doctorData := range doctorsData {
		doctorDto, err := dtos.DoctorDtoFromDoctorDbResponse(doctorData)

		if err != nil {
			return nil, err
		}

		doctors = append(doctors, doctorDto)
	}

	return doctors, err
}

func (service DoctorsService) GetDoctorById(doctorId primitive.ObjectID) (*dtos.DoctorDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getDoctorsCursor(ctx, "_id", doctorId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	doctorData := result[0]
	return dtos.DoctorDtoFromDoctorDbResponse(doctorData)
}

func (service DoctorsService) AddDoctor(doctor models.Doctor) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := hashPassword(doctor.UserPassword)

	if err != nil {
		return err
	}

	doctor.UserPassword = hashedPassword

	coll := configs.GetCollection("doctors")
	res, err := coll.InsertOne(ctx, doctor)

	if err != nil && res == nil {
		return err
	}

	return nil
}

func (service DoctorsService) IsUserIdExist(userId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := bson.D{{"$match", bson.D{{"user_id", userId}}}}

	var err error
	var cursor *mongo.Cursor

	if cursor, err = configs.GetCollection("doctors").Aggregate(ctx, pipeline); err != nil {
		return true
	}

	var result []dbdtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return true
	}

	return len(result) > 0
}

func (service DoctorsService) getDoctorsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
	coll := configs.GetCollection("doctors")

	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "patients",
				"localField":   "_id",
				"foreignField": "_doctor_id",
				"as":           "patients",
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
