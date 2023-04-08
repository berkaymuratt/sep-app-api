package routes

func (routes Routes) defineReportsRoutes() {
	app := routes.app
	controller := routes.reportsController

	doctorsRoutes := app.Group("/api/reports")
	doctorsRoutes.Get("/:id", controller.GetReport)
}
