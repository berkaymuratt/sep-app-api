package routes

func (routes Routes) defineDiseasesRoutes() {
	app := routes.app
	controller := routes.diseasesController

	diseasesRoutes := app.Group("/api/diseases")
	diseasesRoutes.Use(routes.middlewareService.Middleware)
	diseasesRoutes.Get("/", controller.GetDiseases)
	diseasesRoutes.Post("/", controller.AddDisease)
	diseasesRoutes.Put("/:id", controller.UpdateDisease)
	diseasesRoutes.Delete("/:id", controller.DeleteDisease)
}
