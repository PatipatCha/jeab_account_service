package routes

import (
	"github.com/PatipatCha/jeab_ta_service/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	apita := app.Group("/api")

	//
	// *SignIn
	apita.Get("/signin-mobile", controller.SignInForMobile)
	//
	//

}
