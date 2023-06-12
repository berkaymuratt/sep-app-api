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
	"time"
)

var doctorsController controllers.DoctorsController

func TestDoctorsController_GetDoctors(t *testing.T) {
	controller := gomock.NewController(t)

	mockDoctorsService := services.NewMockDoctorsServiceI(controller)
	doctorsController = controllers.NewDoctorsController(mockDoctorsService)

	router := fiber.New()
	router.Get("/api/doctors", doctorsController.GetDoctors)

	doctors := []*dtos.DoctorDto{
		{
			ID:           primitive.NewObjectID(),
			UserId:       "12345678900",
			UserPassword: "password",
			DoctorInfo:   models.DoctorInfo{},
		},
	}

	mockDoctorsService.EXPECT().GetDoctors().Return(doctors, nil)

	target := fmt.Sprintf("/api/doctors")
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestDoctorsController_AddDoctor(t *testing.T) {
	controller := gomock.NewController(t)

	mockDoctorsService := services.NewMockDoctorsServiceI(controller)
	doctorsController = controllers.NewDoctorsController(mockDoctorsService)

	router := fiber.New()
	router.Post("/api/doctors", doctorsController.AddDoctor)

	doctorDto := dtos.DoctorDto{
		ID:           primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "password",
		DoctorInfo:   models.DoctorInfo{},
	}

	newDoctor := models.Doctor{
		UserId:       doctorDto.UserId,
		UserPassword: doctorDto.UserPassword,
		DoctorInfo:   doctorDto.DoctorInfo,
	}

	mockDoctorsService.EXPECT().IsUserIdExist(newDoctor.UserId).Return(false)
	mockDoctorsService.EXPECT().AddDoctor(newDoctor).Return(nil)

	target := fmt.Sprintf("/api/doctors")

	requestBody, _ := json.Marshal(newDoctor)

	req := httptest.NewRequest("POST", target, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	response, _ := router.Test(req)

	assert.Equal(t, 201, response.StatusCode)
}

func TestDoctorsController_GetBusyTimes(t *testing.T) {
	controller := gomock.NewController(t)

	mockDoctorsService := services.NewMockDoctorsServiceI(controller)
	doctorsController = controllers.NewDoctorsController(mockDoctorsService)

	router := fiber.New()
	router.Get("/api/doctors/:id/busy-times", doctorsController.GetBusyTimes)

	doctorId := primitive.NewObjectID()
	date := "2023-12-22T11:30:00.000Z"
	dateTime, _ := time.Parse("2006-01-02T15:04:05.000Z", date)

	var times []time.Time

	mockDoctorsService.EXPECT().GetBusyTimes(doctorId, dateTime).Return(times, nil)

	target := fmt.Sprintf("/api/doctors/%s/busy-times?doctor_id=%s&date=%s", doctorId.Hex(), doctorId, date)
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}
