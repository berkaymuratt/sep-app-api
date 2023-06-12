package tests

import (
	"fmt"
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/dtos"
	services "github.com/berkaymuratt/sep-app-api/mocks/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var authController controllers.AuthController

func TestAuthController_LoginAsPatient(t *testing.T) {
	controller := gomock.NewController(t)

	mockAuthService := services.NewMockAuthServiceI(controller)
	mockJwtService := services.NewMockJwtServiceI(controller)
	mockPatientsService := services.NewMockPatientsServiceI(controller)
	mockDoctorsService := services.NewMockDoctorsServiceI(controller)

	authController = controllers.NewAuthService(mockAuthService, mockJwtService, mockPatientsService, mockDoctorsService)

	router := fiber.New()
	router.Post("/api/auth/login-as-patient", authController.LoginAsPatient)

	userID := "12345678900"
	userPassword := "asdfg"

	patientDto := dtos.PatientDto{}

	mockAuthService.EXPECT().LoginAsPatient(userID, userPassword).Return(&patientDto, nil)
	mockJwtService.EXPECT().GenerateJwtToken(userID).Return("new-token", nil)

	target := fmt.Sprintf("/api/auth/login-as-patient?user_id=%s&user_password=%s", userID, userPassword)
	req := httptest.NewRequest("POST", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestAuthController_LoginAsDoctor(t *testing.T) {
	controller := gomock.NewController(t)

	mockAuthService := services.NewMockAuthServiceI(controller)
	mockJwtService := services.NewMockJwtServiceI(controller)
	mockPatientsService := services.NewMockPatientsServiceI(controller)
	mockDoctorsService := services.NewMockDoctorsServiceI(controller)

	authController = controllers.NewAuthService(mockAuthService, mockJwtService, mockPatientsService, mockDoctorsService)

	router := fiber.New()
	router.Post("/api/auth/login-as-doctor", authController.LoginAsDoctor)

	userID := "12345678900"
	userPassword := "asdfg"

	doctorDto := dtos.DoctorDto{}

	mockAuthService.EXPECT().LoginAsDoctor(userID, userPassword).Return(&doctorDto, nil)
	mockJwtService.EXPECT().GenerateJwtToken(userID).Return("new-token", nil)

	target := fmt.Sprintf("/api/auth/login-as-doctor?user_id=%s&user_password=%s", userID, userPassword)
	req := httptest.NewRequest("POST", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestAuthController_GetCurrentPatientUser(t *testing.T) {
	controller := gomock.NewController(t)

	mockAuthService := services.NewMockAuthServiceI(controller)
	mockJwtService := services.NewMockJwtServiceI(controller)
	mockPatientsService := services.NewMockPatientsServiceI(controller)
	mockDoctorsService := services.NewMockDoctorsServiceI(controller)

	authController = controllers.NewAuthService(mockAuthService, mockJwtService, mockPatientsService, mockDoctorsService)

	router := fiber.New()
	router.Post("/api/auth/current-patient-user", authController.GetCurrentPatientUser)

	userID := "12345678900"
	token := "token"

	patientDto := dtos.PatientDto{}

	mockJwtService.EXPECT().CheckJwt(token).Return(userID, nil)
	mockPatientsService.EXPECT().GetPatientByUserId(userID).Return(&patientDto, nil)

	target := fmt.Sprintf("/api/auth/current-patient-user")
	req := httptest.NewRequest("POST", target, nil)
	req.Header.Add("jwt", token)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}
