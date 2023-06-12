package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/mocks/services"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
)

var diseasesController controllers.DiseaseController

func TestDiseaseController_GetDiseases(t *testing.T) {
	controller := gomock.NewController(t)

	mockDiseaseService := services.NewMockDiseasesServiceI(controller)
	diseasesController = controllers.NewDiseasesController(mockDiseaseService)

	router := fiber.New()
	router.Get("/api/diseases", diseasesController.GetDiseases)

	diseases := []*dtos.DiseaseDto{
		{
			ID:       primitive.NewObjectID(),
			Symptoms: []*dtos.SymptomDto{},
			Name:     "name",
			Details:  "Details",
		},
	}

	mockDiseaseService.EXPECT().GetDiseases().Return(diseases, nil)

	target := fmt.Sprintf("/api/diseases")
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestDiseaseController_AddDisease(t *testing.T) {
	controller := gomock.NewController(t)

	mockDiseaseService := services.NewMockDiseasesServiceI(controller)
	diseasesController = controllers.NewDiseasesController(mockDiseaseService)

	router := fiber.New()
	router.Post("/api/diseases", diseasesController.AddDisease)

	diseaseDto := dtos.DiseaseDto{
		ID: primitive.NewObjectID(),
		Symptoms: []*dtos.SymptomDto{
			{
				ID: primitive.NewObjectID(),
				BodyPart: &models.BodyPart{
					ID:   primitive.NewObjectID(),
					Name: "Body Part",
				},
			},
		},
		Name:    "Name",
		Details: "Details",
	}

	var symptomIds []primitive.ObjectID
	for _, symptom := range diseaseDto.Symptoms {
		symptomIds = append(symptomIds, symptom.ID)
	}

	disease := models.Disease{
		SymptomIds: symptomIds,
		Name:       diseaseDto.Name,
		Details:    diseaseDto.Details,
	}

	mockDiseaseService.EXPECT().AddDisease(disease).Return(nil)

	target := fmt.Sprintf("/api/diseases")

	requestBody, err := json.Marshal(diseaseDto)

	if err != nil {
		t.Fatalf(err.Error())
	}

	req := httptest.NewRequest("POST", target, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	response, _ := router.Test(req)

	assert.Equal(t, 201, response.StatusCode)
}

func TestDiseaseController_UpdateDisease(t *testing.T) {
	controller := gomock.NewController(t)

	mockDiseaseService := services.NewMockDiseasesServiceI(controller)
	diseasesController = controllers.NewDiseasesController(mockDiseaseService)

	router := fiber.New()
	router.Put("/api/diseases/:id", diseasesController.UpdateDisease)

	diseaseId := primitive.NewObjectID()

	diseaseDto := dtos.DiseaseDto{
		ID: diseaseId,
		Symptoms: []*dtos.SymptomDto{
			{
				ID: primitive.NewObjectID(),
				BodyPart: &models.BodyPart{
					ID:   primitive.NewObjectID(),
					Name: "Body Part",
				},
			},
		},
		Name:    "Name",
		Details: "Details",
	}

	var symptomIds []primitive.ObjectID
	for _, symptom := range diseaseDto.Symptoms {
		symptomIds = append(symptomIds, symptom.ID)
	}

	disease := models.Disease{
		SymptomIds: symptomIds,
		Name:       diseaseDto.Name,
		Details:    diseaseDto.Details,
	}

	mockDiseaseService.EXPECT().UpdateDisease(diseaseId, disease).Return(nil)

	target := fmt.Sprintf("/api/diseases/%s", diseaseId.Hex())

	requestBody, err := json.Marshal(diseaseDto)

	if err != nil {
		t.Fatalf(err.Error())
	}

	req := httptest.NewRequest("PUT", target, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}
