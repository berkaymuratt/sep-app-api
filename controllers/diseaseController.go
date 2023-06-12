package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiseaseController struct {
	diseasesService services.DiseasesServiceI
}

func NewDiseasesController(diseasesService services.DiseasesServiceI) DiseaseController {
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

func (controller DiseaseController) AddDisease(ctx *fiber.Ctx) error {
	var request dtos.DiseaseDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "invalid disease data")
	}

	var symptomIds []primitive.ObjectID
	for _, symptom := range request.Symptoms {
		symptomIds = append(symptomIds, symptom.ID)
	}

	disease := models.Disease{
		SymptomIds: symptomIds,
		Name:       request.Name,
		Details:    request.Details,
	}

	if err := controller.diseasesService.AddDisease(disease); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller DiseaseController) UpdateDisease(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	diseaseId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid disease_id")
	}

	var request dtos.DiseaseDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "invalid disease data")
	}

	var symptomIds []primitive.ObjectID
	for _, symptom := range request.Symptoms {
		symptomIds = append(symptomIds, symptom.ID)
	}

	disease := models.Disease{
		SymptomIds: symptomIds,
		Name:       request.Name,
		Details:    request.Details,
	}

	if err := controller.diseasesService.UpdateDisease(diseaseId, disease); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller DiseaseController) DeleteDisease(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	diseaseId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return handleError(ctx, "invalid disease_id")
	}

	if err := controller.diseasesService.DeleteDisease(diseaseId); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
}
