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

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []dbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var patients []*dtos.PatientDto

	for _, data := range result {

		doctor := data.Doctors[0]
		doctorDto := dtos.DoctorDto{
			ID:         doctor.ID,
			UserId:     doctor.UserId,
			DoctorInfo: doctor.DoctorInfo,
		}

		patientDTO := dtos.PatientDto{
			ID:          data.ID,
			DoctorId:    data.DoctorId,
			UserId:      data.UserId,
			PatientInfo: data.PatientInfo,
			Doctor:      &doctorDto,
		}

		patients = append(patients, &patientDTO)
	}

	return patients, err
}

func (service PatientsService) GetPatientsByDoctorId(doctorId primitive.ObjectID) ([]*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		bson.M{
			"$match": bson.M{
				"_doctor_id": doctorId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []dbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var patients []*dtos.PatientDto

	for _, data := range result {
		doctor := data.Doctors[0]
		doctorDto := dtos.DoctorDto{
			ID:         doctor.ID,
			UserId:     doctor.UserId,
			DoctorInfo: doctor.DoctorInfo,
		}

		patientDTO := dtos.PatientDto{
			ID:          data.ID,
			DoctorId:    data.DoctorId,
			UserId:      data.UserId,
			PatientInfo: data.PatientInfo,
			Doctor:      &doctorDto,
		}

		patients = append(patients, &patientDTO)
	}

	return patients, err
}

func (service PatientsService) GetPatientById(patientId primitive.ObjectID) (*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		bson.M{
			"$match": bson.M{
				"_id": patientId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 && len(result[0].Doctors) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	data := result[0]
	doctor := data.Doctors[0]
	doctorDto := dtos.DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	patientDTO := dtos.PatientDto{
		ID:          data.ID,
		DoctorId:    data.DoctorId,
		UserId:      data.UserId,
		PatientInfo: data.PatientInfo,
		Doctor:      &doctorDto,
	}

	return &patientDTO, err
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

	var result []dbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return true
	}

	return len(result) > 0
}
