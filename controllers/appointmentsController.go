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

	appointment, err := controller.appointmentsService.GetAppointment(id)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(appointment)
}
