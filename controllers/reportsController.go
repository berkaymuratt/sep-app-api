package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportsController struct {
	reportsService services.ReportsService
}

func NewReportsController(reportsService services.ReportsService) ReportsController {
	return ReportsController{
		reportsService: reportsService,
	}
}

func (controller ReportsController) GetReport(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	reportId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid report id")
	}

	report, err := controller.reportsService.GetReport(reportId)

	if err != nil {
		return handleError(ctx, "Error is occurred")
	}

	return ctx.Status(fiber.StatusOK).JSON(report)
}
