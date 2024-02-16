package handlers

import (
	shared "testfiber/shared/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginCtrl(c *fiber.Ctx) error {
	//validation
	validate := validator.New()
	body := new(UserLogin)
	c.BodyParser(&body)
	if err := validate.Struct(body); err != nil {
		return c.Status(401).JSON(shared.GlobalErrorHandlerResp{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.SendString("hey")
}

func RegisterCtrl(c *fiber.Ctx) error {
	//validation
	validate := validator.New()
	body := new(UserRegister)
	c.BodyParser(&body)
	if err := validate.Struct(body); err != nil {
		return c.Status(401).JSON(shared.GlobalErrorHandlerResp{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.SendString("done")
}
