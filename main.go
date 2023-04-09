package main

import (
	"github.com/berkaymuratt/sep-app-api/configs"
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/routes"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	configs.ConnectDB()
	app := fiber.New()
	app.Use(logger.New())

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, jwt",
		AllowCredentials: true,
		MaxAge:           3600,
	})

	app.Use(corsMiddleware)

	doctorsService := services.NewDoctorsService()
	patientsService := services.NewPatientsService()
	reportsService := services.NewReportsService()
	symptomsService := services.NewSymptomService()
	bodyPartsService := services.NewBodyPartsService()
	appointmentsService := services.NewAppointmentsService(symptomsService)

	doctorsController := controllers.NewDoctorsController(doctorsService)
	patientsController := controllers.NewPatientsController(patientsService)
	reportsController := controllers.NewReportsController(reportsService)
	appointmentsController := controllers.NewAppointmentsController(appointmentsService)
	symptomsController := controllers.NewSymptomsController(symptomsService)
	bodyPartsController := controllers.NewBodyPartsController(bodyPartsService)

	allRoutes := routes.NewRoutes(
		app,
		doctorsController,
		patientsController,
		reportsController,
		appointmentsController,
		symptomsController,
		bodyPartsController,
	)
	allRoutes.DefineRoutes()

	if err := app.Listen("localhost:8080"); err != nil {
		panic("Error")
	}
}
