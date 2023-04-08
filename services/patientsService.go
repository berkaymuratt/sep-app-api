package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/DbDtos"
	"github.com/berkaymuratt/sep-app-api/Dtos"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PatientsService struct{}

func NewPatientsService() PatientsService {
	return PatientsService{}
}

func (service PatientsService) GetPatients() ([]*Dtos.GetPatientResponse, error) {
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

	var result []DbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var patients []*Dtos.GetPatientResponse

	for _, data := range result {
		patientDTO := Dtos.GetPatientResponse{
			ID:          data.ID,
			DoctorId:    data.DoctorId,
			UserId:      data.UserId,
			PatientInfo: data.PatientInfo,
			Doctor:      data.Doctors[0],
		}

		patients = append(patients, &patientDTO)
	}

	return patients, err
}

func (service PatientsService) GetPatientsByDoctorId(doctorId primitive.ObjectID) ([]*Dtos.GetPatientResponse, error) {
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

	var result []DbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var patients []*Dtos.GetPatientResponse

	for _, data := range result {
		patientDTO := Dtos.GetPatientResponse{
			ID:          data.ID,
			DoctorId:    data.DoctorId,
			UserId:      data.UserId,
			PatientInfo: data.PatientInfo,
		}

		if len(data.Doctors) == 1 {
			patientDTO.Doctor = data.Doctors[0]
		}

		patients = append(patients, &patientDTO)
	}

	return patients, err
}

func (service PatientsService) GetPatientById(patientId primitive.ObjectID) (*Dtos.GetPatientResponse, error) {
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

	var result []*DbDtos.GetPatientDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 && len(result[0].Doctors) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	patientData := result[0]
	patient := Dtos.GetPatientResponse{
		ID:          patientData.ID,
		DoctorId:    patientData.DoctorId,
		UserId:      patientData.UserId,
		PatientInfo: patientData.PatientInfo,
		Doctor:      patientData.Doctors[0],
	}

	return &patient, err
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
