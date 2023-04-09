package routes

import (
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app                    *fiber.App
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
	routes.defineDoctorsRoutes()
	routes.definePatientsRoutes()
	routes.defineReportsRoutes()
	routes.defineAppointmentsRoutes()
	routes.defineSymptomsRoutes()
	routes.defineBodyPartsRoutes()
	routes.defineDiseasesRoutes()
}
