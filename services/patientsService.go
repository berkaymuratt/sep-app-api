package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
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

	var patients []*models.Patient
	if err := cursor.All(context.Background(), &patients); err != nil {
		return nil, err
	}

	var patientsDtos []*dtos.PatientDto

	for _, patient := range patients {
		patientDto := dtos.PatientDto{
			ID:          patient.ID,
			DoctorId:    patient.DoctorId,
			UserId:      patient.UserId,
			PatientInfo: patient.PatientInfo,
		}

		patientsDtos = append(patientsDtos, &patientDto)
	}

	return patientsDtos, err
}

func (service PatientsService) GetPatientsByDoctorId(doctorId primitive.ObjectID) ([]*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getPatientsCursor(ctx, "_doctor_id", doctorId)

	if err != nil {
		return nil, err
	}

	var patients []*models.Patient
	if err := cursor.All(context.Background(), &patients); err != nil {
		return nil, err
	}

	var patientsDtos []*dtos.PatientDto

	for _, patient := range patients {
		patientDto := dtos.PatientDto{
			ID:          patient.ID,
			DoctorId:    patient.DoctorId,
			UserId:      patient.UserId,
			PatientInfo: patient.PatientInfo,
		}

		patientsDtos = append(patientsDtos, &patientDto)
	}

	return patientsDtos, err
}

func (service PatientsService) GetPatientById(patientId primitive.ObjectID) (*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getPatientsCursor(ctx, "_id", patientId)

	if err != nil {
		return nil, err
	}

	var patients []*models.Patient
	if err := cursor.All(context.Background(), &patients); err != nil {
		return nil, err
	}

	if len(patients) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	patient := patients[0]
	patientDto := dtos.PatientDto{
		ID:          patient.ID,
		DoctorId:    patient.DoctorId,
		UserId:      patient.UserId,
		PatientInfo: patient.PatientInfo,
	}
	return &patientDto, nil
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

	var result []models.Patient
	if err := cursor.All(context.Background(), &result); err != nil {
		return true
	}

	return len(result) > 0
}

func (service PatientsService) getPatientsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
	coll := configs.GetCollection("patients")
	pipeline := bson.A{}

	if matchField != "" {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				matchField: matchValue,
			},
		})
	}

	return coll.Aggregate(ctx, pipeline)
}
