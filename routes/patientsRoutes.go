package routes

func (routes Routes) definePatientsRoutes() {
	app := routes.app
	controller := routes.patientsController

	patientsRoutes := app.Group("/api/patients")
	patientsRoutes.Get("/", controller.GetPatients)
	patientsRoutes.Get("/:id", controller.GetPatientById)
	patientsRoutes.Post("/", controller.AddPatient)
}
