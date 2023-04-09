package controllers

import (
	"github.com/berkaymuratt/sep-app-api/dtos"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SymptomsController struct {
	symptomsService services.SymptomsService
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
