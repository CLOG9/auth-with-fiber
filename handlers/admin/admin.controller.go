package admin_controller

import "github.com/gofiber/fiber/v2"

func AdminCtrl(c *fiber.Ctx) error {
	return c.SendString("hey")
}
