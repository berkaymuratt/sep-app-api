package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbDtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DoctorsService struct{}

func NewDoctorsService() DoctorsService {
	return DoctorsService{}
}

func (service DoctorsService) GetDoctors() ([]*dtos.DoctorDto, error) {
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

	var doctorsData []*dbDtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &doctorsData); err != nil {
		return nil, err
	}

	var doctors []*dtos.DoctorDto

	for _, doctorData := range doctorsData {

		var patientsData []*dtos.PatientDto

		for _, patient := range doctorData.Patients {
			patientDto := dtos.PatientDto{
				ID:          patient.ID,
				DoctorId:    patient.DoctorId,
				UserId:      patient.UserId,
				PatientInfo: patient.PatientInfo,
			}
			patientsData = append(patientsData, &patientDto)
		}

		doctor := dtos.DoctorDto{
			ID:         doctorData.ID,
			UserId:     doctorData.UserId,
			DoctorInfo: doctorData.DoctorInfo,
			Patients:   patientsData,
		}
		doctors = append(doctors, &doctor)
	}

	return doctors, err
}

func (service DoctorsService) GetDoctorById(doctorId primitive.ObjectID) (*dtos.DoctorDto, error) {
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

	var result []*dbDtos.GetDoctorDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	doctorData := result[0]

	var patientsData []*dtos.PatientDto

	for _, patient := range doctorData.Patients {
		patientDto := dtos.PatientDto{
			ID:          patient.ID,
			DoctorId:    patient.DoctorId,
			UserId:      patient.UserId,
			PatientInfo: patient.PatientInfo,
		}
		patientsData = append(patientsData, &patientDto)
	}

	doctor := dtos.DoctorDto{
		ID:         doctorData.ID,
		UserId:     doctorData.UserId,
		DoctorInfo: doctorData.DoctorInfo,
		Patients:   patientsData,
	}
	return &doctor, err
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
