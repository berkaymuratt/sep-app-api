package controllers

import (
	models "github.com/berkaymuratt/sep-app-api/models/doctor"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoctorsController struct {
	doctorsService services.DoctorsService
}

func NewDoctorsController(doctorsService services.DoctorsService) DoctorsController {
	return DoctorsController{
		doctorsService: doctorsService,
	}
}

func (controller DoctorsController) GetDoctors(ctx *fiber.Ctx) error {
	doctorsService := controller.doctorsService
	doctors, err := doctorsService.GetDoctors()

	if err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"doctors": doctors,
	})
}

func (controller DoctorsController) GetDoctorById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	doctorId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "patient_id is invalid")
	}

	doctorsService := controller.doctorsService
	patient, err := doctorsService.GetDoctorById(doctorId)

	if err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusOK).JSON(patient)
}

func (controller DoctorsController) AddDoctor(ctx *fiber.Ctx) error {
	doctorsService := controller.doctorsService

	var request models.AddDoctorRequest
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "Invalid Patient Data")
	}

	newDoctor := models.Doctor{
		UserId:       request.UserId,
		UserPassword: request.UserPassword,
		DoctorInfo:   request.DoctorInfo,
	}

	if err := doctorsService.AddDoctor(newDoctor); err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
	})
}
