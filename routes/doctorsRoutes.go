package routes

func (routes Routes) defineDoctorsRoutes() {
	app := routes.app
	controller := routes.doctorsController

	doctorsRoutes := app.Group("/api/doctors")
	doctorsRoutes.Use(routes.middlewareService.Middleware)
	doctorsRoutes.Get("/", controller.GetDoctors)
	doctorsRoutes.Get("/:id", controller.GetDoctorById)
	doctorsRoutes.Patch("/:id", controller.UpdateDoctor)
	doctorsRoutes.Post("/", controller.AddDoctor)
	doctorsRoutes.Get("/:id/busy-times", controller.GetBusyTimes)
}
