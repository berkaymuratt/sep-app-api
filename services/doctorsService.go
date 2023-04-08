package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/DbDtos"
	"github.com/berkaymuratt/sep-app-api/Dtos"
	"github.com/berkaymuratt/sep-app-api/configs"
	doctor_models "github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DoctorsService struct{}

func NewDoctorsService() DoctorsService {
	return DoctorsService{}
}

func (service DoctorsService) GetDoctors() ([]*Dtos.GetDoctorResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var doctorsData []*DbDtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &doctorsData); err != nil {
		return nil, err
	}

	var doctors []*Dtos.GetDoctorResponse
	for _, doctorData := range doctorsData {
		doctor := Dtos.GetDoctorResponse{
			ID:         doctorData.ID,
			UserId:     doctorData.UserId,
			DoctorInfo: doctorData.DoctorInfo,
			Patients:   doctorData.Patients,
		}
		doctors = append(doctors, &doctor)
	}

	return doctors, err
}

func (service DoctorsService) GetDoctorById(doctorId primitive.ObjectID) (*Dtos.GetDoctorResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		bson.M{
			"$match": bson.M{
				"_id": doctorId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*DbDtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	doctorData := result[0]
	doctor := Dtos.GetDoctorResponse{
		ID:         doctorData.ID,
		UserId:     doctorData.UserId,
		DoctorInfo: doctorData.DoctorInfo,
		Patients:   doctorData.Patients,
	}
	return &doctor, err
}

func (service DoctorsService) AddDoctor(doctor doctor_models.Doctor) error {
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
