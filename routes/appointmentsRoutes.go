package routes

func (routes Routes) defineAppointmentsRoutes() {
	app := routes.app
	controller := routes.appointmentsController

	appointmentsRoutes := app.Group("/api/appointments")
	appointmentsRoutes.Use(routes.middlewareService.Middleware)
	appointmentsRoutes.Get("/", controller.GetAppointments)
	appointmentsRoutes.Post("/", controller.AddAppointment)
	appointmentsRoutes.Get("/:id", controller.GetAppointmentById)
	appointmentsRoutes.Delete("/:id", controller.DeleteAppointmentById)
	appointmentsRoutes.Patch("/:id/date", controller.UpdateAppointmentDate)
}
