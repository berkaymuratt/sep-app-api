package routes

func (routes Routes) definePatientsRoutes() {
	app := routes.app
	controller := routes.patientsController

	patientsRoutes := app.Group("/api/patients")
	patientsRoutes.Use(routes.middlewareService.Middleware)
	patientsRoutes.Get("/", controller.GetPatients)
	patientsRoutes.Get("/:id", controller.GetPatientById)
	patientsRoutes.Patch("/:id", controller.UpdatePatient)
	patientsRoutes.Post("/", controller.AddPatient)
}
