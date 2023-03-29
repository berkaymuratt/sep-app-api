package services

import (
	"context"
	"github.com/berkaymuratt/sep-app-api/config"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type DoctorsService struct{}

func NewDoctorsService() DoctorsService {
	return DoctorsService{}
}

func (service DoctorsService) GetDoctors() ([]*models.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("doctors")
	cursor, err := coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var doctors []*models.Doctor
	if err := cursor.All(context.Background(), &doctors); err != nil {
		log.Fatal(err)
	}

	return doctors, err
}
