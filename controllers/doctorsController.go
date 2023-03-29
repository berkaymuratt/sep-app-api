package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
)

type DoctorsController struct {
	doctorsService services.DoctorsService
}

func NewDoctorsController(doctorsService services.DoctorsService) DoctorsController {
	return DoctorsController{
		doctorsService: doctorsService,
	}
}

func (controller DoctorsController) GetDoctors(ctx *fiber.Ctx) error {
	doctorsService := controller.doctorsService
	doctors, err := doctorsService.GetDoctors()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error is occurred",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"doctors": doctors,
	})
}
