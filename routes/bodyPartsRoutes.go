package routes

func (routes Routes) defineBodyPartsRoutes() {
	app := routes.app
	controller := routes.bodyPartsController

	bodyPartsRoutes := app.Group("/api/body-parts")
	bodyPartsRoutes.Use(routes.middlewareService.Middleware)
	bodyPartsRoutes.Get("/", controller.GetBodyParts)
}
