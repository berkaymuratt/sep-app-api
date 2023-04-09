package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppointmentsController struct {
	appointmentsService services.AppointmentsService
	reportsService      services.ReportsService
}

func NewAppointmentsController(appointmentsService services.AppointmentsService, reportsService services.ReportsService) AppointmentsController {
	return AppointmentsController{
		appointmentsService: appointmentsService,
		reportsService:      reportsService,
	}
}

func (controller AppointmentsController) GetAppointmentById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid appointment_id")
	}

	appointment, err := controller.appointmentsService.GetAppointmentById(id)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(appointment)
}

func (controller AppointmentsController) GetAppointments(ctx *fiber.Ctx) error {
	idStr := ctx.Query("doctor_id")
	doctorId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid doctor_id")
	}

	appointments, err := controller.appointmentsService.GetAppointmentByDoctor(doctorId)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"appointments": appointments,
	})
}

func (controller AppointmentsController) AddAppointment(ctx *fiber.Ctx) error {
	appointmentsService := controller.appointmentsService
	reportsService := controller.reportsService

	var request dtos.AppointmentDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "invalid appointment data")
	}

	if !appointmentsService.IsDateAvailable(request.Doctor.ID, request.Patient.ID, request.Date) {
		return handleError(ctx, "date is not available")
	}

	var symptomIds []primitive.ObjectID

	for _, symptomDto := range request.Symptoms {
		symptomIds = append(symptomIds, symptomDto.ID)
	}

	appointment := models.Appointment{
		ID:          request.ID,
		DoctorId:    request.Doctor.ID,
		PatientId:   request.Patient.ID,
		SymptomIds:  symptomIds,
		PatientNote: request.PatientNote,
		Date:        request.Date,
	}

	if err := reportsService.CreateReport(&appointment); err != nil {
		return handleError(ctx, err.Error())
	}

	if err := appointmentsService.AddAppointment(appointment); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
	})
}
