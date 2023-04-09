package routes

func (routes Routes) defineAppointmentsRoutes() {
	app := routes.app
	controller := routes.appointmentsController

	appointmentsRoutes := app.Group("/api/appointments")
	appointmentsRoutes.Get("/:id", controller.GetAppointmentById)

}
