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
	doctorIdStr := ctx.Query("doctor_id")
	patientIdStr := ctx.Query("patient_id")

	var appointments []*dtos.AppointmentDto
	var err error

	if doctorIdStr != "" {
		doctorId, err := primitive.ObjectIDFromHex(doctorIdStr)

		if err != nil {
			return handleError(ctx, "invalid doctor_id")
		}

		appointments, err = controller.appointmentsService.GetAppointments(&doctorId, nil)

	} else if patientIdStr != "" {
		patientId, err := primitive.ObjectIDFromHex(patientIdStr)

		if err != nil {
			return handleError(ctx, "invalid patient_id")
		}

		appointments, err = controller.appointmentsService.GetAppointments(nil, &patientId)
	} else {
		return handleError(ctx, "missing id")
	}

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

	isAvailable := appointmentsService.IsDateAvailable(request.Doctor.ID, request.Patient.ID, request.Date)

	if !isAvailable {
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

	if err := reportsService.CreateReportByAppointment(&appointment); err != nil {
		return handleError(ctx, err.Error())
	}

	if err := appointmentsService.AddAppointment(appointment); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller AppointmentsController) UpdateAppointmentDate(ctx *fiber.Ctx) error {
	appointmentsService := controller.appointmentsService

	idStr := ctx.Params("id")
	appointmentId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid appointment_id")
	}

	var request dtos.AppointmentDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "invalid appointment data")
	}

	var symptomIds []primitive.ObjectID
	for _, symptomDto := range request.Symptoms {
		symptomIds = append(symptomIds, symptomDto.ID)
	}

	updatedAppointment := models.Appointment{
		ID:          request.ID,
		DoctorId:    request.Doctor.ID,
		PatientId:   request.Patient.ID,
		SymptomIds:  symptomIds,
		PatientNote: request.PatientNote,
		Date:        request.Date,
	}

	isAvailable := appointmentsService.IsDateAvailable(updatedAppointment.DoctorId, updatedAppointment.PatientId, updatedAppointment.Date)

	if !isAvailable {
		return handleError(ctx, "invalid date")
	}

	err = controller.appointmentsService.UpdateAppointmentDate(appointmentId, updatedAppointment.Date)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller AppointmentsController) DeleteAppointmentById(ctx *fiber.Ctx) error {

	var err error

	idStr := ctx.Params("id")
	appointmentId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid appointment_id")
	}

	err = controller.appointmentsService.DeleteAppointment(appointmentId)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
}
