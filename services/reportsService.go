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

//go:generate mockgen -destination=../mocks/service/mockReportsService.go -package=services github.com/berkaymuratt/sep-app-api/services ReportsServiceI
type ReportsServiceI interface {
	GetReportById(reportId primitive.ObjectID) (*dtos.ReportDto, error)
	GetReports(doctorId *primitive.ObjectID, patientId *primitive.ObjectID) ([]*dtos.ReportDto, error)
	CreateReportByAppointment(appointment *models.Appointment) error
	DeleteReport(reportId primitive.ObjectID) error
	AddDoctorFeedback(reportId primitive.ObjectID, feedback string) error
}

type ReportsService struct {
	ReportsServiceI
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

func (service ReportsService) GetReports(doctorId *primitive.ObjectID, patientId *primitive.ObjectID) ([]*dtos.ReportDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cursor *mongo.Cursor
	var err error

	if doctorId != nil {
		cursor, err = service.getReportsCursor(ctx, "_doctor_id", doctorId)
	} else if patientId != nil {
		cursor, err = service.getReportsCursor(ctx, "_patient_id", patientId)
	} else {
		return nil, errors.New("missing ids")
	}

	if err != nil {
		return nil, err
	}

	var result []*dbdtos.GetReportDbResponse
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	var reportsDtos []*dtos.ReportDto

	for _, reportData := range result {
		var symptomsDto []*dtos.SymptomDto
		if symptomsDto, err = service.symptomsService.GetSymptomsByIds(reportData.SymptomIds); err != nil {
			return nil, err
		}

		var diseasesDto []*dtos.DiseaseDto
		if diseasesDto, err = service.diseasesService.GetDiseasesByIds(reportData.PossibleDiseaseIds); err != nil {
			return nil, err
		}

		reportDto, err := dtos.ReportDtoFromReportDbResponse(reportData, symptomsDto, diseasesDto)

		if err != nil {
			return nil, err
		}

		reportsDtos = append(reportsDtos, reportDto)
	}

	return reportsDtos, err
}

func (service ReportsService) AddDoctorFeedback(reportId primitive.ObjectID, feedback string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := configs.GetCollection("reports")

	update := bson.M{
		"$set": bson.M{
			"doctor_feedback": feedback,
		},
	}

	_, err := coll.UpdateByID(ctx, reportId, update)
	return err
}

func (service ReportsService) CreateReportByAppointment(appointment *models.Appointment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	coll := configs.GetCollection("diseases")

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"_symptom_ids": bson.M{
					"$all": appointment.SymptomIds,
				},
			},
		},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return err
	}

	var diseases []*models.Disease
	if err := cursor.All(context.Background(), &diseases); err != nil {
		return err
	}

	var possibleDiseaseIds []primitive.ObjectID
	for _, disease := range diseases {
		possibleDiseaseIds = append(possibleDiseaseIds, disease.ID)
	}

	var diseaseIds []primitive.ObjectID

	if possibleDiseaseIds != nil {
		diseaseIds = possibleDiseaseIds
	}

	report := models.Report{
		DoctorId:           appointment.DoctorId,
		PatientId:          appointment.PatientId,
		SymptomIds:         appointment.SymptomIds,
		PossibleDiseaseIds: diseaseIds,
		DoctorFeedback:     "",
		PatientNote:        appointment.PatientNote,
		CreatedAt:          time.Now(),
	}

	res, err := configs.GetCollection("reports").InsertOne(ctx, report)

	if err != nil {
		return err
	}

	appointment.ReportId = res.InsertedID.(primitive.ObjectID)
	return err
}

func (service ReportsService) DeleteReport(reportId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	coll := configs.GetCollection("reports")
	_, err := coll.DeleteOne(ctx, bson.M{"_id": reportId})

	return err
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
