package routes

func (routes Routes) defineBodyPartsRoutes() {
	app := routes.app
	controller := routes.bodyPartsController

	bodyPartsRoutes := app.Group("/api/body-parts")
	bodyPartsRoutes.Get("/", controller.GetBodyParts)
}
