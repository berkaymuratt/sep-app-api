package controllers

import (
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	authService     services.AuthService
	jwtService      services.JwtService
	patientsService services.PatientsService
	doctorsService  services.DoctorsService
}

func NewAuthService(authService services.AuthService, jwtService services.JwtService) AuthController {
	return AuthController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (controller AuthController) UpdateDoctorPassword(ctx *fiber.Ctx) error {
	doctorId := ctx.Query("doctor_id")
	newPassword := ctx.Query("new_password")

	id, err := primitive.ObjectIDFromHex(doctorId)

	if err != nil {
		return handleError(ctx, err.Error())
	}

	if err := controller.authService.UpdateDoctorPassword(id, newPassword); err != nil {
		return handleError(ctx, "password could not be updated")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful",
	})
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
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"patient":   &patient,
		"sep-token": token,
	})
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
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"doctor":    &doctor,
		"sep-token": token,
	})
}

func (controller AuthController) GetCurrentPatientUser(ctx *fiber.Ctx) error {
	patientsService := controller.patientsService
	jwtService := controller.jwtService

	jwt := ctx.Get("jwt")

	userId, err := jwtService.CheckJwt(jwt)

	if err != nil {
		return handleErrorWithStatus(ctx, fiber.StatusUnauthorized, "patient id cannot found..")
	}

	patient, err := patientsService.GetPatientByUserId(userId)

	if err != nil {
		return handleErrorWithStatus(ctx, fiber.StatusUnauthorized, "patient cannot found..")
	}

	return ctx.Status(fiber.StatusOK).JSON(&patient)
}

func (controller AuthController) GetCurrentDoctorUser(ctx *fiber.Ctx) error {
	doctorsService := controller.doctorsService
	jwtService := controller.jwtService

	jwt := ctx.Get("jwt")

	userId, err := jwtService.CheckJwt(jwt)

	if err != nil {
		return handleErrorWithStatus(ctx, fiber.StatusUnauthorized, "doctor id cannot found..")
	}

	if err != nil {
		return handleError(ctx, "error is occured")
	}

	doctor, err := doctorsService.GetDoctorByUserId(userId)

	if err != nil {
		return handleErrorWithStatus(ctx, fiber.StatusUnauthorized, "doctor cannot found..")
	}

	return ctx.Status(fiber.StatusOK).JSON(&doctor)
}
