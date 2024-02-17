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
	// Validation
	validate := validator.New()

	// Parse request body into UserRegister struct
	body := new(UserRegister)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(shared.GlobErrResp(err.Error()))
	}

	// Validate request body
	if err := validate.Struct(body); err != nil {
		return c.Status(401).JSON(shared.GlobErrResp(err.Error()))
	}
	if err := createUser(body); err != nil {
		return c.Status(400).JSON(shared.GlobErrResp(err.Error()))
	}
	return c.JSON(shared.GlobResp(body))
}
