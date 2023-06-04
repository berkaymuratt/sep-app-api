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

	var doctors []*models.Doctor
	if err := cursor.All(context.Background(), &doctors); err != nil {
		return nil, err
	}

	var doctorsDtos []*dtos.DoctorDto

	for _, doctor := range doctors {
		doctorDto := dtos.DoctorDto{
			ID:         doctor.ID,
			UserId:     doctor.UserId,
			DoctorInfo: doctor.DoctorInfo,
		}

		doctorsDtos = append(doctorsDtos, &doctorDto)
	}

	return doctorsDtos, err
}

func (service DoctorsService) GetDoctorById(doctorId primitive.ObjectID) (*dtos.DoctorDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getDoctorsCursor(ctx, "_id", doctorId)

	if err != nil {
		return nil, err
	}

	var result []*models.Doctor
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	doctor := result[0]

	doctorDto := dtos.DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	return &doctorDto, nil
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

func (service DoctorsService) GetDoctorByUserId(userId string) (*dtos.DoctorDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getDoctorsCursor(ctx, "user_id", userId)

	if err != nil {
		return nil, err
	}

	var result []*models.Doctor
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 {
		return nil, errors.New("doctor_models cannot found")
	}

	doctor := result[0]

	doctorDto := dtos.DoctorDto{
		ID:         doctor.ID,
		UserId:     doctor.UserId,
		DoctorInfo: doctor.DoctorInfo,
	}

	return &doctorDto, nil
}

func (service DoctorsService) UpdateDoctor(doctorId primitive.ObjectID, newDoctorDto dtos.DoctorDto) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("doctors")

	update := bson.M{
		"$set": bson.M{
			"doctor_info": newDoctorDto.DoctorInfo,
		},
	}

	_, err := coll.UpdateByID(ctx, doctorId, update)
	return err
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

	var result []*models.Doctor
	if err := cursor.All(context.Background(), &result); err != nil {
		return true
	}

	return len(result) > 0
}

func (service DoctorsService) GetBusyTimes(doctorId primitive.ObjectID, date time.Time) ([]time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	year := date.Year()
	month := date.Month()
	day := date.Day()

	filter := bson.M{
		"_doctor_id": doctorId,
		"date": bson.M{
			"$gte": time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
			"$lt":  time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC),
		},
	}

	var err error
	var cursor *mongo.Cursor

	if cursor, err = configs.GetCollection("appointments").Find(ctx, filter); err != nil {
		return nil, err
	}

	var result []*dtos.AppointmentDto
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var appointmentTimes []time.Time

	for _, result := range result {
		appointmentTimes = append(appointmentTimes, result.Date)
	}

	return appointmentTimes, nil
}

func (service DoctorsService) getDoctorsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
	coll := configs.GetCollection("doctors")

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
