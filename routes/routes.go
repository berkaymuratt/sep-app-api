package routes

import (
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/berkaymuratt/sep-app-api/services"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app                    *fiber.App
	middlewareService      services.MiddlewareService
	authController         controllers.AuthController
	doctorsController      controllers.DoctorsController
	patientsController     controllers.PatientsController
	reportsController      controllers.ReportsController
	appointmentsController controllers.AppointmentsController
	symptomsController     controllers.SymptomsController
	bodyPartsController    controllers.BodyPartsController
	diseasesController     controllers.DiseaseController
}

func NewRoutes(
	app *fiber.App,
	middlewareService services.MiddlewareService,
	authController controllers.AuthController,
	doctorsController controllers.DoctorsController,
	patientsController controllers.PatientsController,
	reportsController controllers.ReportsController,
	appointmentsController controllers.AppointmentsController,
	symptomsController controllers.SymptomsController,
	bodyPartsController controllers.BodyPartsController,
	diseasesController controllers.DiseaseController,
) Routes {
	return Routes{
		app:                    app,
		middlewareService:      middlewareService,
		authController:         authController,
		doctorsController:      doctorsController,
		patientsController:     patientsController,
		reportsController:      reportsController,
		appointmentsController: appointmentsController,
		symptomsController:     symptomsController,
		bodyPartsController:    bodyPartsController,
		diseasesController:     diseasesController,
	}
}

func (routes Routes) DefineRoutes() {
	routes.defineAuthRoutes()
	routes.defineDoctorsRoutes()
	routes.definePatientsRoutes()
	routes.defineReportsRoutes()
	routes.defineAppointmentsRoutes()
	routes.defineSymptomsRoutes()
	routes.defineBodyPartsRoutes()
	routes.defineDiseasesRoutes()
}
