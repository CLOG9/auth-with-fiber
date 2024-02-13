package main

import (
	"testfiber/config"
	admin_route "testfiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api := app.Group(config.Defaults.ApiVersion) // api
	admin_route.Admin_route(api)

	//
	app.Listen(":3000")
}
