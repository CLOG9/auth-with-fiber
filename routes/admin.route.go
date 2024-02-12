package admin_route

import (
	"testfiber/config"

	"github.com/gofiber/fiber/v2"
)

func Admin_route(api fiber.Router) {
	admin := api.Group(config.RouteEndpts.Admin) // admin

	admin.Get("/state", func(c *fiber.Ctx) error {
		return c.SendString("admin dashboard")
	})

}
