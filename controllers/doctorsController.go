package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DoctorsController struct {
	doctorsService services.DoctorsServiceI
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
		return handleError(ctx, "invalid doctor_id")
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

	var request dtos.DoctorDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "Invalid Doctor Data")
	}

	if doctorsService.IsUserIdExist(request.UserId) {
		return handleError(ctx, "user_id has already exist in the system")
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

func (controller DoctorsController) UpdateDoctor(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	doctorId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid doctor_id")
	}

	var doctorDto dtos.DoctorDto
	if err := ctx.BodyParser(&doctorDto); err != nil {
		return handleError(ctx, "invalid doctor data")
	}

	if err := controller.doctorsService.UpdateDoctor(doctorId, doctorDto); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller DoctorsController) GetBusyTimes(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	doctorId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid doctor_id")
	}

	date := ctx.Query("date")
	dateTime, err := time.Parse("2006-01-02T15:04:05.000Z", date)

	if err != nil {
		return handleError(ctx, "invalid date")
	}

	times, err := controller.doctorsService.GetBusyTimes(doctorId, dateTime)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"busy_times": times,
	})
}
