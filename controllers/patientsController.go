package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientsController struct {
	patientsService services.PatientsService
}

func NewPatientsController(patientsService services.PatientsService) PatientsController {
	return PatientsController{
		patientsService: patientsService,
	}
}

func (controller PatientsController) GetPatients(ctx *fiber.Ctx) error {
	patientsService := controller.patientsService

	var patients []*dtos.PatientDto
	var err error

	var doctorId primitive.ObjectID
	idStr := ctx.Query("doctor_id")

	if idStr != "" {
		doctorId, err = primitive.ObjectIDFromHex(idStr)

		if err != nil {
			return handleError(ctx, "doctor_id is invalid")
		}

		patients, err = patientsService.GetPatientsByDoctorId(doctorId)
	} else {
		patients, err = patientsService.GetPatients()
	}

	if err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"patients": patients,
	})
}

func (controller PatientsController) GetPatientById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	patientId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "patient_id is invalid")
	}

	patientsService := controller.patientsService
	patient, err := patientsService.GetPatientById(patientId)

	if err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusOK).JSON(patient)
}

func (controller PatientsController) AddPatient(ctx *fiber.Ctx) error {
	patientsService := controller.patientsService

	var request dtos.PatientDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "Invalid Patient Data")
	}

	newPatient := models.Patient{
		DoctorId:     request.DoctorId,
		UserId:       request.UserId,
		UserPassword: request.UserPassword,
		PatientInfo:  request.PatientInfo,
	}

	if err := patientsService.AddPatient(newPatient); err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
	})
}
