package controllers

import "github.com/gofiber/fiber/v2"

func handleError(c *fiber.Ctx, errMessage string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": errMessage,
	})
}

func handleErrorWithStatus(c *fiber.Ctx, statusCode int, errMessage string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": errMessage,
	})
}
