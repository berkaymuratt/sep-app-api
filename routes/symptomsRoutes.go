package routes

func (routes Routes) defineSymptomsRoutes() {
	app := routes.app
	controller := routes.symptomsController

	symptomsRoutes := app.Group("/api/symptoms")
	symptomsRoutes.Get("/", controller.GetSymptoms)
}
