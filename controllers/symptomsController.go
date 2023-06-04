package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/models"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SymptomsController struct {
	symptomsService services.SymptomsServiceI
}

func NewSymptomsController(symptomsService services.SymptomsService) SymptomsController {
	return SymptomsController{
		symptomsService: symptomsService,
	}
}

func (controller SymptomsController) GetSymptoms(ctx *fiber.Ctx) error {
	symptomsService := controller.symptomsService

	var symptoms []*dtos.SymptomDto
	var err error

	idStr := ctx.Query("body_part_id")
	bodyPartId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil || idStr == "" {
		symptoms, err = symptomsService.GetSymptoms()
	} else {
		symptoms, err = symptomsService.GetSymptomsByBodyPart(bodyPartId)
	}

	if err != nil {
		return handleError(ctx, "symptoms cannot found")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"symptoms": symptoms,
	})
}

func (controller SymptomsController) AddSymptom(ctx *fiber.Ctx) error {
	symptomsService := controller.symptomsService

	var request *dtos.SymptomDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "invalid symptom data")
	}

	symptom := models.Symptom{
		ID:         request.ID,
		BodyPartId: request.BodyPart.ID,
		Name:       request.Name,
		Level:      request.Level,
	}

	if err := symptomsService.AddSymptom(symptom); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller SymptomsController) UpdateSymptom(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	symptomId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return err
	}

	var request *dtos.SymptomDto
	if err := ctx.BodyParser(&request); err != nil {
		return handleError(ctx, "invalid symptom data")
	}

	symptom := models.Symptom{
		ID:         request.ID,
		BodyPartId: request.BodyPart.ID,
		Name:       request.Name,
		Level:      request.Level,
	}

	if err := controller.symptomsService.UpdateSymptom(symptomId, symptom); err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
}

func (controller SymptomsController) DeleteSymptom(ctx *fiber.Ctx) error {

	var err error

	idStr := ctx.Params("id")
	symptomId, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return err
	}

	err = controller.symptomsService.DeleteSymptom(symptomId)

	if err != nil {
		return handleError(ctx, err.Error())
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "successful",
		})
	}
}
