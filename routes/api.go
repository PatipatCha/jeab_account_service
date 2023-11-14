package routes

import (
	"github.com/PatipatCha/jeab_ta_service/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// *JGuard
	jguard := v1.Group("/jguard")
	jguard.Post("/signin", controller.SignInAndChangeMobileHandler)
	jguard.Post("/validate-otp", controller.VaildOTPHandler)
	jguard.Put("/pdpa", controller.UpdatePDPAHandler)
	jguard.Put("/profile", controller.UpdateProfileHandler)
	jguard.Put("/change-mobile", controller.SignInAndChangeMobileHandler)

	// *JMaster
	jmaster := v1.Group("/jmaster")
	jmaster.Post("/signin", controller.SignInForJMaster)

	// *JCenter
	jcenter := v1.Group("/jcenter")
	jcenter.Post("/signin", controller.SignInForJCenter)

}
