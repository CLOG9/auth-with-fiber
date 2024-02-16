package handlers

import "github.com/gofiber/fiber/v2"

func AdminCtrl(c *fiber.Ctx) error {
	return c.SendString("hey")
}
