package routes

func (routes Routes) defineDiseasesRoutes() {
	app := routes.app
	controller := routes.diseasesController

	diseasesRoutes := app.Group("/api/diseases")
	diseasesRoutes.Get("/", controller.GetDiseases)
}
