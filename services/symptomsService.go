package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SymptomsService struct{}

func NewSymptomService() SymptomsService {
	return SymptomsService{}
}

func (service SymptomsService) GetSymptoms() ([]*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("symptoms")
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "body_parts",
				"localField":   "_body_part_id",
				"foreignField": "_id",
				"as":           "body_parts",
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetSymptomDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var symptomsDtos []*dtos.SymptomDto
	for _, symptom := range result {
		symptomDto := dtos.SymptomDto{
			ID:       symptom.ID,
			BodyPart: &symptom.BodyParts[0],
			Name:     symptom.Name,
			Level:    symptom.Level,
		}
		symptomsDtos = append(symptomsDtos, &symptomDto)
	}

	return symptomsDtos, nil
}

func (service SymptomsService) GetSymptomById(symptomId primitive.ObjectID) (*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("symptoms")
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "body_parts",
				"localField":   "_body_part_id",
				"foreignField": "_id",
				"as":           "body_parts",
			},
		},
		bson.M{
			"$match": bson.M{
				"_id": symptomId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetSymptomDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) != 1 && len(result[0].BodyParts) != 1 {
		return nil, errors.New("symptom cannot found")
	}

	data := result[0]
	symptomDto := dtos.SymptomDto{
		ID:       data.ID,
		BodyPart: &data.BodyParts[0],
	}

	return &symptomDto, nil
}

func (service SymptomsService) GetSymptomsByIds(symptomIds []primitive.ObjectID) ([]dbdtos.GetSymptomDbResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("symptoms")

	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "body_parts",
				"localField":   "_body_part_id",
				"foreignField": "_id",
				"as":           "body_parts",
			},
		},
		bson.M{
			"$match": bson.M{
				"_id": bson.M{
					"$in": symptomIds,
				},
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []dbdtos.GetSymptomDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (service SymptomsService) GetSymptomsByBodyPart(bodyPartId primitive.ObjectID) ([]*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("symptoms")
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "body_parts",
				"localField":   "_body_part_id",
				"foreignField": "_id",
				"as":           "body_parts",
			},
		},
		bson.M{
			"$match": bson.M{
				"_body_part_id": bodyPartId,
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetSymptomDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var symptomsDtos []*dtos.SymptomDto
	for _, symptom := range result {
		symptomDto := dtos.SymptomDto{
			ID:       symptom.ID,
			BodyPart: &symptom.BodyParts[0],
			Name:     symptom.Name,
			Level:    symptom.Level,
		}
		symptomsDtos = append(symptomsDtos, &symptomDto)
	}

	return symptomsDtos, nil
}
