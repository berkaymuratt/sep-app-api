package routes

import (
	"github.com/berkaymuratt/sep-app-api/controllers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app               *fiber.App
	doctorsController controllers.DoctorsController
}

func NewRoutes(app *fiber.App, doctorsController controllers.DoctorsController) Routes {
	return Routes{
		app:               app,
		doctorsController: doctorsController,
	}
}

func (routes Routes) DefineRoutes() {
	routes.defineDoctorsRoutes()
}

func (routes Routes) defineDoctorsRoutes() {
	app := routes.app
	controller := routes.doctorsController

	doctorsRoutes := app.Group("/api/doctors")
	doctorsRoutes.Get("/", controller.GetDoctors)
}
