package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
)

type BodyPartsController struct {
	bodyPartsService services.BodyPartsServiceI
}

func NewBodyPartsController(bodyPartsService services.BodyPartsServiceI) BodyPartsController {
	return BodyPartsController{
		bodyPartsService: bodyPartsService,
	}
}

func (controller BodyPartsController) GetBodyParts(ctx *fiber.Ctx) error {
	bodyPartsService := controller.bodyPartsService

	bodyParts, err := bodyPartsService.GetBodyParts()

	if err != nil {
		return handleError(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"body_parts": bodyParts,
	})
}
