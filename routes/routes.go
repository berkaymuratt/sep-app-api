package routes

import (
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app                *fiber.App
	doctorsController  controllers.DoctorsController
	patientsController controllers.PatientsController
	reportsController  controllers.ReportsController
}

func NewRoutes(app *fiber.App, doctorsController controllers.DoctorsController, patientsController controllers.PatientsController, reportsController controllers.ReportsController) Routes {
	return Routes{
		app:                app,
		doctorsController:  doctorsController,
		patientsController: patientsController,
		reportsController:  reportsController,
	}
}

func (routes Routes) DefineRoutes() {
	routes.defineDoctorsRoutes()
	routes.definePatientsRoutes()
	routes.defineReportsRoutes()
}
