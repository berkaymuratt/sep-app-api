package routes

func (routes Routes) defineAuthRoutes() {
	app := routes.app
	controller := routes.authController

	authRoutes := app.Group("/api/auth")
	authRoutes.Post("/login-as-doctor", controller.LoginAsDoctor)
	authRoutes.Post("/login-as-patient", controller.LoginAsPatient)
	authRoutes.Get("/current-patient-user", controller.GetCurrentPatientUser)
	authRoutes.Get("/current-doctor-user", controller.GetCurrentDoctorUser)
	authRoutes.Patch("/update-doctor-password", controller.UpdateDoctorPassword)
}
