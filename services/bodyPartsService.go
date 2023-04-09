package services

import (
	"context"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type BodyPartsService struct{}

func NewBodyPartsService() BodyPartsService {
	return BodyPartsService{}
}

func (service BodyPartsService) GetBodyParts() ([]*models.BodyPart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("body_parts")
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var bodyParts []*models.BodyPart
	if err := cursor.All(context.Background(), &bodyParts); err != nil {
		return nil, err
	}

	return bodyParts, nil
}
