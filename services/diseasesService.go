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
	"time"
)

type DiseasesService struct {
	symptomsService SymptomsService
}

func NewDiseasesService(symptomsService SymptomsService) DiseasesService {
	return DiseasesService{
		symptomsService: symptomsService,
	}
}

func (service DiseasesService) GetDiseases() ([]*dtos.DiseaseDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("diseases")
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetDiseaseDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var diseasesDtos []*dtos.DiseaseDto

	for _, diseasesData := range result {
		var symptomsData []*dtos.SymptomDto
		if symptomsData, err = service.symptomsService.GetSymptomsByIds(diseasesData.SymptomIds); err != nil {
			return nil, err
		}

		diseasesDto, err := dtos.DiseaseDtoFromDiseaseDbResponse(diseasesData, symptomsData)

		if err != nil {
			return nil, err
		}

		diseasesDtos = append(diseasesDtos, diseasesDto)
	}

	return diseasesDtos, nil
}

func (service DiseasesService) GetDiseasesByIds(diseasesIds []primitive.ObjectID) ([]*dtos.DiseaseDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("diseases")

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"_id": bson.M{"$in": diseasesIds},
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetDiseaseDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var diseasesDtos []*dtos.DiseaseDto

	for _, diseasesData := range result {
		var symptomsData []*dtos.SymptomDto
		if symptomsData, err = service.symptomsService.GetSymptomsByIds(diseasesData.SymptomIds); err != nil {
			return nil, err
		}

		diseasesDto, err := dtos.DiseaseDtoFromDiseaseDbResponse(diseasesData, symptomsData)

		if err != nil {
			return nil, err
		}

		diseasesDtos = append(diseasesDtos, diseasesDto)
	}

	return diseasesDtos, nil
}

func (service DiseasesService) AddDisease(newDisease models.Disease) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("diseases")
	_, err := coll.InsertOne(ctx, newDisease)
	return err
}

func (service DiseasesService) DeleteDisease(diseaseId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reportsColl := configs.GetCollection("reports")

	res, err := reportsColl.Find(ctx, bson.M{
		"_possible_disease_ids": diseaseId,
	})

	if err != nil {
		return err
	}

	var report []*models.Report
	if err := res.All(ctx, &report); err != nil {
		return err
	}

	if len(report) > 0 {
		return errors.New("disease is in report(s)")
	}

	coll := configs.GetCollection("diseases")
	_, err = coll.DeleteOne(ctx, bson.M{
		"_id": diseaseId,
	})
	return err
}

func (service DiseasesService) UpdateDisease(diseaseId primitive.ObjectID, updatedDisease models.Disease) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updatedDisease,
	}

	_, err := configs.GetCollection("diseases").UpdateByID(ctx, diseaseId, update)
	return err
}
