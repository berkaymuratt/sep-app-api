package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
)

type DiseaseController struct {
	diseasesService services.DiseasesService
}

func NewDiseasesController(diseasesService services.DiseasesService) DiseaseController {
	return DiseaseController{
		diseasesService: diseasesService,
	}
}

func (controller DiseaseController) GetDiseases(ctx *fiber.Ctx) error {
	diseases, err := controller.diseasesService.GetDiseases()

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"diseases": diseases,
	})
}
