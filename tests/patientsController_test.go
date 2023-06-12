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

var patientsController controllers.PatientsController

func TestPatientsController_GetPatients(t *testing.T) {
	controller := gomock.NewController(t)

	mockPatientsService := services.NewMockPatientsServiceI(controller)
	patientsController = controllers.NewPatientsController(mockPatientsService)

	router := fiber.New()
	router.Get("/api/patients", patientsController.GetPatients)

	patients := []*dtos.PatientDto{
		{
			ID:           primitive.NewObjectID(),
			UserId:       "12345678900",
			UserPassword: "password",
			PatientInfo:  models.PatientInfo{},
		},
	}

	mockPatientsService.EXPECT().GetPatients().Return(patients, nil)

	target := fmt.Sprintf("/api/patients")
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestPatientsController_GetPatientById(t *testing.T) {
	controller := gomock.NewController(t)

	mockPatientsService := services.NewMockPatientsServiceI(controller)
	patientsController = controllers.NewPatientsController(mockPatientsService)

	patientId := primitive.NewObjectID()

	router := fiber.New()
	router.Get("/api/patients/:id", patientsController.GetPatientById)

	patient := dtos.PatientDto{
		ID:           patientId,
		UserId:       "12345678900",
		UserPassword: "password",
		PatientInfo:  models.PatientInfo{},
	}

	mockPatientsService.EXPECT().GetPatientById(patientId).Return(&patient, nil)

	target := fmt.Sprintf("/api/patients/%s", patientId.Hex())
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestPatientsController_AddPatient(t *testing.T) {
	controller := gomock.NewController(t)

	mockPatientsService := services.NewMockPatientsServiceI(controller)
	patientsController = controllers.NewPatientsController(mockPatientsService)

	router := fiber.New()
	router.Post("/api/patients", patientsController.AddPatient)

	patientDto := dtos.PatientDto{
		ID:           primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "password",
		PatientInfo:  models.PatientInfo{},
	}

	newPatient := models.Patient{
		UserId:       patientDto.UserId,
		UserPassword: patientDto.UserPassword,
		PatientInfo:  patientDto.PatientInfo,
	}

	mockPatientsService.EXPECT().IsUserIdExist(newPatient.UserId).Return(false)
	mockPatientsService.EXPECT().AddPatient(newPatient).Return(nil)

	target := fmt.Sprintf("/api/patients")

	requestBody, _ := json.Marshal(newPatient)

	req := httptest.NewRequest("POST", target, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	response, _ := router.Test(req)

	assert.Equal(t, 201, response.StatusCode)
}
