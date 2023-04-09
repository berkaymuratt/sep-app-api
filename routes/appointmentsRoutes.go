package routes

func (routes Routes) defineAppointmentsRoutes() {
	app := routes.app
	controller := routes.appointmentsController

	doctorsRoutes := app.Group("/api/appointments")
	doctorsRoutes.Get("/:id", controller.GetAppointmentById)

}
