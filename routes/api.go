package routes

import (
	"github.com/PatipatCha/jeab_ta_service/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	//
	// *SignIn Mobile
	api.Post("/v1/signin", controller.SignIn)
	api.Post("/v1/validate-otp", controller.VaildOTP)
	api.Post("/v1/pdpa", controller.UpdatePDPA)
	//
	//

	// *JGuard
	// SignIn
	jguard := v1.Group("/jguard")
	jguard.Post("/signin", controller.SignIn)
	jguard.Post("/validate-otp", controller.VaildOTP)
	//
	// PDPA
	jguard.Post("/pdpa", controller.UpdatePDPA)
	//
	//

	// *JMaster
	// **SignIn
	jmaster := v1.Group("/jmaster")
	jmaster.Post("/signin", controller.SignInForJMaster)
	jmaster.Get("/pdpa", controller.GetPDPA)
	//
	//

	// *Profile
	// profile := v1.Group("/profile")
	// profile.Post("/", controller.GetProfile)
	//

}
