package routes

func (routes Routes) defineReportsRoutes() {
	app := routes.app
	controller := routes.reportsController

	reportsRoutes := app.Group("/api/reports")
	reportsRoutes.Get("/:id", controller.GetReport)
}
