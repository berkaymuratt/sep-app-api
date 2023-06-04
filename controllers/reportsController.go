package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportsController struct {
	reportsService services.ReportsServiceI
}

func NewReportsController(reportsService services.ReportsServiceI) ReportsController {
	return ReportsController{
		reportsService: reportsService,
	}
}

func (controller ReportsController) GetReportById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	reportId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid report id")
	}

	report, err := controller.reportsService.GetReportById(reportId)

	if err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusOK).JSON(report)
}

func (controller ReportsController) GetReports(ctx *fiber.Ctx) error {
	doctorIdStr := ctx.Query("doctor_id")
	patientIdStr := ctx.Query("patient_id")

	var reports []*dtos.ReportDto
	var err error

	if doctorIdStr != "" {
		doctorId, err := primitive.ObjectIDFromHex(doctorIdStr)

		if err != nil {
			return handleError(ctx, "invalid doctor_id")
		}

		reports, err = controller.reportsService.GetReports(&doctorId, nil)

	} else if patientIdStr != "" {
		patientId, err := primitive.ObjectIDFromHex(patientIdStr)

		if err != nil {
			return handleError(ctx, "invalid patient_id")
		}

		reports, err = controller.reportsService.GetReports(nil, &patientId)
	} else {
		return handleError(ctx, "missing id")
	}

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"reports": reports,
	})
}
