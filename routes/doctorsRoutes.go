package routes

func (routes Routes) defineDoctorsRoutes() {
	app := routes.app
	controller := routes.doctorsController

	doctorsRoutes := app.Group("/api/doctors")
	doctorsRoutes.Get("/", controller.GetDoctors)
	doctorsRoutes.Get("/:id", controller.GetDoctorById)
	doctorsRoutes.Patch("/:id", controller.UpdateDoctor)
	doctorsRoutes.Post("/", controller.AddDoctor)
}
