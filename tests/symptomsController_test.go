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

var symptomsController controllers.SymptomsController

func TestSymptomsController_GetSymptoms(t *testing.T) {
	controller := gomock.NewController(t)

	mockSymptomsService := services.NewMockSymptomsServiceI(controller)
	symptomsController = controllers.NewSymptomsController(mockSymptomsService)

	router := fiber.New()
	router.Get("/api/symptoms", symptomsController.GetSymptoms)

	bodyPartId := primitive.NewObjectID()

	symptoms := []*dtos.SymptomDto{
		{},
	}

	mockSymptomsService.EXPECT().GetSymptomsByBodyPart(bodyPartId).Return(symptoms, nil)

	target := fmt.Sprintf("/api/symptoms?body_part_id=%s", bodyPartId.Hex())
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestSymptomsController_AddSymptom(t *testing.T) {
	controller := gomock.NewController(t)

	mockSymptomsService := services.NewMockSymptomsServiceI(controller)
	symptomsController = controllers.NewSymptomsController(mockSymptomsService)

	router := fiber.New()
	router.Post("/api/symptoms", symptomsController.AddSymptom)

	symptomDto := dtos.SymptomDto{
		ID: primitive.NewObjectID(),
		BodyPart: &models.BodyPart{
			ID:   primitive.NewObjectID(),
			Name: "Body Part",
		},
		Name:  "Symptom",
		Level: 0,
	}

	symptom := models.Symptom{
		ID:         symptomDto.ID,
		BodyPartId: symptomDto.BodyPart.ID,
		Name:       symptomDto.Name,
		Level:      symptomDto.Level,
	}

	mockSymptomsService.EXPECT().AddSymptom(symptom).Return(nil)

	requestBody, _ := json.Marshal(symptomDto)

	target := fmt.Sprintf("/api/symptoms")
	req := httptest.NewRequest("POST", target, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	response, _ := router.Test(req)

	assert.Equal(t, 201, response.StatusCode)
}

func TestSymptomsController_DeleteSymptom(t *testing.T) {
	controller := gomock.NewController(t)

	mockSymptomsService := services.NewMockSymptomsServiceI(controller)
	symptomsController = controllers.NewSymptomsController(mockSymptomsService)

	router := fiber.New()
	router.Delete("/api/symptoms/:id", symptomsController.DeleteSymptom)

	symptomId := primitive.NewObjectID()

	mockSymptomsService.EXPECT().DeleteSymptom(symptomId).Return(nil)

	target := fmt.Sprintf("/api/symptoms/%s", symptomId.Hex())
	req := httptest.NewRequest("DELETE", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}
