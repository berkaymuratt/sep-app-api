package routes

func (routes Routes) defineSymptomsRoutes() {
	app := routes.app
	controller := routes.symptomsController

	symptomsRoutes := app.Group("/api/symptoms")
	symptomsRoutes.Use(routes.middlewareService.Middleware)
	symptomsRoutes.Get("/", controller.GetSymptoms)
	symptomsRoutes.Post("/", controller.AddSymptom)
	symptomsRoutes.Delete("/:id", controller.DeleteSymptom)
	symptomsRoutes.Put("/:id", controller.UpdateSymptom)
}
