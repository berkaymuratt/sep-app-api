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

type ReportsService struct {
	symptomsService SymptomsService
	diseasesService DiseasesService
}

func NewReportsService() ReportsService {
	return ReportsService{}
}

func (service ReportsService) GetReportById(reportId primitive.ObjectID) (*dtos.ReportDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := service.getReportsCursor(ctx, "_id", reportId)

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetReportDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("cannot found any report")
	}

	if len(result) != 1 && len(result[0].Doctors) != 1 {
		return nil, errors.New("doctor cannot found")
	}

	if len(result) != 1 && len(result[0].Patients) != 1 {
		return nil, errors.New("patient cannot found")
	}

	reportData := result[0]

	symptomsDtos, err := service.symptomsService.GetSymptomsByIds(reportData.SymptomIds)

	if err != nil {
		return nil, err
	}

	diseasesDtos, err := service.diseasesService.GetDiseasesByIds(reportData.PossibleDiseaseIds)

	if err != nil {
		return nil, err
	}

	return dtos.ReportDtoFromReportDbResponse(reportData, symptomsDtos, diseasesDtos)
}

func (service ReportsService) getReportsCursor(ctx context.Context, matchField string, matchValue any) (*mongo.Cursor, error) {
	coll := configs.GetCollection("reports")
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
			"$lookup": bson.M{
				"from":         "patients",
				"localField":   "_patient_id",
				"foreignField": "_id",
				"as":           "patients",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "symptoms",
				"localField":   "_symptom_ids",
				"foreignField": "_id",
				"as":           "symptoms",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "diseases",
				"localField":   "_possible_disease_ids",
				"foreignField": "_id",
				"as":           "possible_diseases",
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
