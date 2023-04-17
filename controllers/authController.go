package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.AuthService
	jwtService  services.JwtService
}

func NewAuthService(authService services.AuthService, jwtService services.JwtService) AuthController {
	return AuthController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (controller AuthController) LoginAsPatient(ctx *fiber.Ctx) error {
	userId := ctx.Query("user_id")
	userPassword := ctx.Query("user_password")

	patient, err := controller.authService.LoginAsPatient(userId, userPassword)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	token, err := controller.jwtService.GenerateJwtToken(userId)

	if err != nil {
		return err
	}

	ctx.Set("jwt", token)
	return ctx.Status(fiber.StatusOK).JSON(&patient)
}

func (controller AuthController) LoginAsDoctor(ctx *fiber.Ctx) error {
	userId := ctx.Query("user_id")
	userPassword := ctx.Query("user_password")

	doctor, err := controller.authService.LoginAsDoctor(userId, userPassword)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	token, err := controller.jwtService.GenerateJwtToken(userId)

	if err != nil {
		return err
	}

	ctx.Set("jwt", token)
	return ctx.Status(fiber.StatusOK).JSON(&doctor)
}
