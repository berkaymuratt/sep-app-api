package routes

func (routes Routes) defineReportsRoutes() {
	app := routes.app
	controller := routes.reportsController

	reportsRoutes := app.Group("/api/reports")
	reportsRoutes.Use(routes.middlewareService.Middleware)
	reportsRoutes.Get("/", controller.GetReports)
	reportsRoutes.Get("/:id", controller.GetReportById)
}
