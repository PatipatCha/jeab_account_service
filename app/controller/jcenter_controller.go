package controller

import (
	"os"
	"strings"

	"github.com/PatipatCha/jeab_ta_service/app/model"
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func SignInForJCenter(c *fiber.Ctx) error {
	var rawBody = model.WebSignInRequest{}
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	userId := strings.TrimSpace(rawBody.UserId)
	passcode := strings.TrimSpace(rawBody.Passcode)

	isUser := services.FindUser(userId, "seoc")
	if !isUser {
		res := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("CHECK_USER_NOT_FOUND"),
		}
		return c.Status(404).JSON(res)
	}

	// check user passcode
	user, _ := services.CheckUserPasscode(userId, passcode)
	if user.Firstname == "" {
		res := fiber.Map{
			"user_id": userId,
			"data":    fiber.Map{},
			"message": os.Getenv("SIGN_IN_WRONG"),
		}
		return c.Status(401).JSON(res)
	}

	res := fiber.Map{
		"user_id": userId,
		"data":    user,
		"message": "SIGN IN SUCCESS",
	}

	return c.Status(200).JSON(res)
}
