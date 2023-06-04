package services

import (
	"context"
	"errors"
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/dbdtos"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//go:generate mockgen -destination=../mocks/service/mockSymptomsService.go -package=services github.com/berkaymuratt/sep-app-api/services SymptomsServiceI
type SymptomsServiceI interface {
	GetSymptoms() ([]*dtos.SymptomDto, error)
	GetSymptomById(symptomId primitive.ObjectID) (*dtos.SymptomDto, error)
	GetSymptomsByIds(symptomIds []primitive.ObjectID) ([]*dtos.SymptomDto, error)
	GetSymptomsByBodyPart(bodyPartId primitive.ObjectID) ([]*dtos.SymptomDto, error)
	AddSymptom(newSymptom models.Symptom) error
	UpdateSymptom(symptomId primitive.ObjectID, updatedSymptom models.Symptom) error
	DeleteSymptom(symptomId primitive.ObjectID) error
}

type SymptomsService struct {
	SymptomsServiceI
}

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

func (service SymptomsService) AddSymptom(newSymptom models.Symptom) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := configs.GetCollection("symptoms").InsertOne(ctx, newSymptom)
	return err
}

func (service SymptomsService) UpdateSymptom(symptomId primitive.ObjectID, updatedSymptom models.Symptom) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updatedSymptom,
	}

	_, err := configs.GetCollection("symptoms").UpdateByID(ctx, symptomId, update)
	return err
}

func (service SymptomsService) DeleteSymptom(symptomId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var res *mongo.Cursor
	var err error

	reportsColl := configs.GetCollection("reports")
	res, err = reportsColl.Find(ctx, bson.M{
		"_symptom_ids": symptomId,
	})

	if err != nil {
		return err
	}

	var foundFromReports []*models.Report
	if err := res.All(ctx, &foundFromReports); err != nil {
		return err
	}

	diseasesColl := configs.GetCollection("diseases")
	res, err = diseasesColl.Find(ctx, bson.M{
		"_symptom_ids": symptomId,
	})

	if err != nil {
		return err
	}

	var foundFromDiseases []*models.Report
	if err := res.All(ctx, &foundFromDiseases); err != nil {
		return err
	}

	if len(foundFromReports) > 0 || len(foundFromDiseases) > 0 {
		return errors.New("symptom id is used in other report(s) or/and disease(s)")
	}

	coll := configs.GetCollection("symptoms")
	_, err = coll.DeleteOne(ctx, bson.M{"_id": symptomId})
	return err
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
