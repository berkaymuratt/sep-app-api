package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type SymptomsService struct{}

func NewSymptomService() SymptomsService {
	return SymptomsService{}
}

func (service SymptomsService) GetSymptoms() ([]*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getSymptomsCursor(ctx, "", "")

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetSymptomDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var symptomsDtos []*dtos.SymptomDto
	for _, symptomData := range result {
		symptomDto, err := dtos.SymptomDtoFromSymptomDbResponse(symptomData)

		if err != nil {
			return nil, err
		}

		symptomsDtos = append(symptomsDtos, symptomDto)
	}

	return symptomsDtos, nil
}

func (service SymptomsService) GetSymptomById(symptomId primitive.ObjectID) (*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getSymptomsCursor(ctx, "_id", symptomId)

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

	symptomData := result[0]
	return dtos.SymptomDtoFromSymptomDbResponse(symptomData)
}

func (service SymptomsService) GetSymptomsByIds(symptomIds []primitive.ObjectID) ([]*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getSymptomsCursor(ctx, "_id", bson.M{"$in": symptomIds})

	if err != nil {
		return nil, err
	}

	var result []dbdtos.GetSymptomDbResponse
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

func (service SymptomsService) GetSymptomsByBodyPart(bodyPartId primitive.ObjectID) ([]*dtos.SymptomDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.getSymptomsCursor(ctx, "_body_part_id", bodyPartId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetSymptomDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var symptomsDtos []*dtos.SymptomDto
	for _, symptomData := range result {
		symptomDto, err := dtos.SymptomDtoFromSymptomDbResponse(symptomData)

		if err != nil {
			return nil, err
		}

		symptomsDtos = append(symptomsDtos, symptomDto)
	}

	return symptomsDtos, nil
}

func (service SymptomsService) getSymptomsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
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

	if matchField != "" {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				matchField: matchValue,
			},
		})
	}

	return coll.Aggregate(ctx, pipeline)
}
