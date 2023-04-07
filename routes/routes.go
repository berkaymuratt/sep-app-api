package routes

import (
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app                *fiber.App
	doctorsController  controllers.DoctorsController
	patientsController controllers.PatientsController
}

func NewRoutes(app *fiber.App, doctorsController controllers.DoctorsController, patientsController controllers.PatientsController) Routes {
	return Routes{
		app:                app,
		doctorsController:  doctorsController,
		patientsController: patientsController,
	}
}

func (routes Routes) DefineRoutes() {
	routes.defineDoctorsRoutes()
	routes.definePatientsRoutes()
}
