package routes

import (
	"github.com/PatipatCha/jeab_ta_service/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	api := app.Group("/api")

	//
	// *SignIn Mobile
	api.Post("/v1/signin", controller.SignIn)
	api.Post("/v1/validate-otp", controller.VaildOTP)
	api.Post("/v1/update-pdpa", controller.UpdatePDPA)
	//
	// *SignIn Web
	api.Post("/v1/signin-jmaster", controller.SignInForJMaster)
	//

	// *Profile
	api.Post("/v1/get-profile", controller.GetProfile)
	//

}
