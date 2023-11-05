package controller

import (
	"os"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func SignInForMobile(c *fiber.Ctx) error {
	var data = model.MobileOTPSignInResponse{}
	mobileNumber := c.Query("mobile")
	if mobileNumber == "" {
		output := fiber.Map{
			"Mobile":  mobileNumber,
			"Data":    fiber.Map{},
			"Message": os.Getenv("MOBILE_IS_NULL"),
		}
		return c.JSON(output)
	}
	users, _ := services.CheckUserId(mobileNumber)

	if users.UserId == "" {
		output := fiber.Map{
			"Mobile":  mobileNumber,
			"Data":    fiber.Map{},
			"Message": os.Getenv("CHECK_MOBILE_NOT_FOUND"),
		}
		return c.JSON(output)
	} else {
		data, _ = services.SendOTPServer(mobileNumber)
	}

	res := model.MobileSignInResponse{
		Mobile:  mobileNumber,
		Data:    data,
		Message: os.Getenv("CHECK_MOBILE_INPROCESS"),
	}

	return c.JSON(res)
}

// func MobileNumberVerification(c *fiber.Ctx) error {

// }
