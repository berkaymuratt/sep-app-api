package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/dtos"
	services "github.com/berkaymuratt/sep-app-api/mocks/services"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
	"time"
)

var appointmentsController controllers.AppointmentsController

func TestAppointmentsController_GetAppointments(t *testing.T) {
	controller := gomock.NewController(t)

	mockAppointmentsService := services.NewMockAppointmentsServiceI(controller)
	mockReportsService := services.NewMockReportsServiceI(controller)

	appointmentsController = controllers.NewAppointmentsController(mockAppointmentsService, mockReportsService)

	router := fiber.New()
	router.Get("/api/appointments", appointmentsController.GetAppointments)

	fakeDoctor := dtos.DoctorDto{
		ID:           primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "hashedPassword",
		DoctorInfo: models.DoctorInfo{
			Name:      "DoctorName",
			Surname:   "DoctorSurname",
			Address:   "Address",
			Telephone: "0123456789",
		},
	}

	fakePatient := dtos.PatientDto{
		ID:           primitive.NewObjectID(),
		DoctorId:     primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "hashedPassword",
		PatientInfo: models.PatientInfo{
			Name:      "PatientName",
			Surname:   "PatientSurname",
			Gender:    "Gender",
			Age:       20,
			Height:    175,
			Weight:    70,
			Address:   "Address",
			Telephone: "0123456789",
		},
	}

	fakeSymptoms :=
		[]*dtos.SymptomDto{
			{
				ID: primitive.NewObjectID(),
				BodyPart: &models.BodyPart{
					ID:   primitive.NewObjectID(),
					Name: "BodyPartName",
				},
				Name:  "SymptomName",
				Level: 0,
			},
		}

	fakeAppointments := []*dtos.AppointmentDto{
		{
			ID:          primitive.NewObjectID(),
			ReportId:    primitive.NewObjectID(),
			Doctor:      &fakeDoctor,
			Patient:     &fakePatient,
			Symptoms:    fakeSymptoms,
			PatientNote: "PatientNote",
			Date:        time.Now(),
		},
	}

	doctorID := primitive.NewObjectID()

	mockAppointmentsService.EXPECT().GetAppointments(&doctorID, nil).Return(fakeAppointments, nil)

	target := fmt.Sprintf("/api/appointments?doctor_id=%s", doctorID.Hex())
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestAppointmentsController_GetAppointmentById(t *testing.T) {
	controller := gomock.NewController(t)

	mockAppointmentsService := services.NewMockAppointmentsServiceI(controller)
	mockReportsService := services.NewMockReportsServiceI(controller)

	appointmentsController = controllers.NewAppointmentsController(mockAppointmentsService, mockReportsService)

	router := fiber.New()
	router.Get("/api/appointments/:id", appointmentsController.GetAppointmentById)

	fakeDoctor := dtos.DoctorDto{
		ID:           primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "hashedPassword",
		DoctorInfo: models.DoctorInfo{
			Name:      "DoctorName",
			Surname:   "DoctorSurname",
			Address:   "Address",
			Telephone: "0123456789",
		},
	}

	fakePatient := dtos.PatientDto{
		ID:           primitive.NewObjectID(),
		DoctorId:     primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "hashedPassword",
		PatientInfo: models.PatientInfo{
			Name:      "PatientName",
			Surname:   "PatientSurname",
			Gender:    "Gender",
			Age:       20,
			Height:    175,
			Weight:    70,
			Address:   "Address",
			Telephone: "0123456789",
		},
	}

	fakeSymptoms :=
		[]*dtos.SymptomDto{
			{
				ID: primitive.NewObjectID(),
				BodyPart: &models.BodyPart{
					ID:   primitive.NewObjectID(),
					Name: "BodyPartName",
				},
				Name:  "SymptomName",
				Level: 0,
			},
		}

	fakeAppointment := dtos.AppointmentDto{
		ID:          primitive.NewObjectID(),
		ReportId:    primitive.NewObjectID(),
		Doctor:      &fakeDoctor,
		Patient:     &fakePatient,
		Symptoms:    fakeSymptoms,
		PatientNote: "PatientNote",
		Date:        time.Now(),
	}

	appointmentID := primitive.NewObjectID()

	mockAppointmentsService.EXPECT().GetAppointmentById(appointmentID).Return(&fakeAppointment, nil)

	target := fmt.Sprintf("/api/appointments/%s", appointmentID.Hex())
	req := httptest.NewRequest("GET", target, nil)
	response, _ := router.Test(req)

	assert.Equal(t, 200, response.StatusCode)
}

func TestAppointmentsController_AddAppointment(t *testing.T) {
	controller := gomock.NewController(t)

	mockAppointmentsService := services.NewMockAppointmentsServiceI(controller)
	mockReportsService := services.NewMockReportsServiceI(controller)

	appointmentsController = controllers.NewAppointmentsController(mockAppointmentsService, mockReportsService)

	router := fiber.New()
	router.Post("/api/appointments", appointmentsController.AddAppointment)

	fakeDoctor := dtos.DoctorDto{
		ID:           primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "hashedPassword",
		DoctorInfo: models.DoctorInfo{
			Name:      "DoctorName",
			Surname:   "DoctorSurname",
			Address:   "Address",
			Telephone: "0123456789",
		},
	}

	fakePatient := dtos.PatientDto{
		ID:           primitive.NewObjectID(),
		DoctorId:     primitive.NewObjectID(),
		UserId:       "12345678900",
		UserPassword: "hashedPassword",
		PatientInfo: models.PatientInfo{
			Name:      "PatientName",
			Surname:   "PatientSurname",
			Gender:    "Gender",
			Age:       20,
			Height:    175,
			Weight:    70,
			Address:   "Address",
			Telephone: "0123456789",
		},
	}

	fakeSymptoms :=
		[]*dtos.SymptomDto{
			{
				ID: primitive.NewObjectID(),
				BodyPart: &models.BodyPart{
					ID:   primitive.NewObjectID(),
					Name: "BodyPartName",
				},
				Name:  "SymptomName",
				Level: 0,
			},
		}

	fakeAppointmentDto := dtos.AppointmentDto{
		Doctor:      &fakeDoctor,
		Patient:     &fakePatient,
		Symptoms:    fakeSymptoms,
		PatientNote: "PatientNote",
		Date:        time.Time{},
	}

	var symptomIds []primitive.ObjectID
	for _, symptomDto := range fakeAppointmentDto.Symptoms {
		symptomIds = append(symptomIds, symptomDto.ID)
	}

	fakeAppointment := models.Appointment{
		ID:          fakeAppointmentDto.ID,
		DoctorId:    fakeAppointmentDto.Doctor.ID,
		PatientId:   fakeAppointmentDto.Patient.ID,
		SymptomIds:  symptomIds,
		PatientNote: fakeAppointmentDto.PatientNote,
		Date:        fakeAppointmentDto.Date,
	}

	mockAppointmentsService.EXPECT().IsDateAvailable(fakeAppointmentDto.Doctor.ID, fakeAppointmentDto.Patient.ID, fakeAppointmentDto.Date).Return(true)
	mockReportsService.EXPECT().CreateReportByAppointment(&fakeAppointment).Return(nil)
	mockAppointmentsService.EXPECT().AddAppointment(fakeAppointment).Return(nil)

	target := "/api/appointments"

	requestBody, err := json.Marshal(fakeAppointmentDto)

	if err != nil {
		t.Fatalf(err.Error())
	}

	req := httptest.NewRequest("POST", target, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	response, _ := router.Test(req)
	assert.Equal(t, 201, response.StatusCode)
}
