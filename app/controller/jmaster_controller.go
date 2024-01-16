package controller

import (
	"os"
	"strings"

	"github.com/PatipatCha/jeab_account_service/app/model"
	"github.com/PatipatCha/jeab_account_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func LogInForJMaster(c *fiber.Ctx) error {
	var rawBody = model.JMasterLogInRequest{}
	if err := c.BodyParser(&rawBody); err != nil {
		return err
	}

	userId := strings.TrimSpace(rawBody.Username)
	passcode := strings.TrimSpace(rawBody.Password)

	user, _ := services.CheckUserJMasterPassword(userId, passcode)
	if user.JeabID == "" {
		res := fiber.Map{
			"data":    fiber.Map{},
			"message": os.Getenv("LOGIN_INCORRECT"),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	res := fiber.Map{
		"data":    user,
		"message": "LOGIN SUCCESS",
	}

	return c.JSON(res)
}
