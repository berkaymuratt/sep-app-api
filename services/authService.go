package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (service AuthService) LoginAsDoctor(userId string, userPassword string) (*dtos.DoctorDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("doctors")
	res := coll.FindOne(ctx, bson.M{
		"user_id": userId,
	})

	var doctor *models.Doctor
	if err := res.Decode(&doctor); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(doctor.UserPassword), []byte(userPassword)); err != nil {
		return nil, errors.New("username or password is wrong")
	}

	doctorDto := dtos.DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	return &doctorDto, nil
}

func (service AuthService) LoginAsPatient(userId string, userPassword string) (*dtos.PatientDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("patients")
	res := coll.FindOne(ctx, bson.M{
		"user_id": userId,
	})

	var patient *models.Patient
	if err := res.Decode(&patient); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(patient.UserPassword), []byte(userPassword)); err != nil {
		return nil, errors.New("username or password is wrong")
	}

	patientDto := dtos.PatientDto{
		ID:          patient.ID,
		DoctorId:    patient.DoctorId,
		UserId:      patient.UserId,
		PatientInfo: patient.PatientInfo,
	}

	return &patientDto, nil
}

func (service AuthService) UpdatePatientPassword(patientId primitive.ObjectID, newPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := hashPassword(newPassword)

	if err != nil {
		return err
	}

	coll := configs.GetCollection("patients")

	update := bson.M{
		"$set": bson.M{
			"user_password": hashedPassword,
		},
	}

	_, err = coll.UpdateByID(ctx, patientId, update)
	return err
}

func (service AuthService) UpdateDoctorPassword(doctorId primitive.ObjectID, newPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := hashPassword(newPassword)

	if err != nil {
		return err
	}

	coll := configs.GetCollection("doctors")

	update := bson.M{
		"$set": bson.M{
			"user_password": hashedPassword,
		},
	}

	_, err = coll.UpdateByID(ctx, doctorId, update)
	return err
}
