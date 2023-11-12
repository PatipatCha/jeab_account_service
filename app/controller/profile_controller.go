package controller

import (
	"github.com/PatipatCha/jeab_ta_service/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetPDPA(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	user, _ := services.GetUser("", userId, "")
	res := fiber.Map{
		"user_id": userId,
		"data": fiber.Map{
			"personal_pdpa":        user.PersonalPDPA,
			"personal_expire_date": user.PersonalExpireDate,
		},
		"message": "PDPA Lists",
	}
	return c.JSON(res)
}
