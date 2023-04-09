package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppointmentsController struct {
	appointmentsService services.AppointmentsService
}

func NewAppointmentsController(appointmentsService services.AppointmentsService) AppointmentsController {
	return AppointmentsController{
		appointmentsService: appointmentsService,
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
