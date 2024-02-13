package main

import (
	"fmt"
	"testfiber/config"
	say_hey "testfiber/middlewares"
	admin_route "testfiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(say_hey.SayHey)
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("hey mom")
		return c.SendString("Hello, World!")
	})
	api := app.Group(config.Defaults.ApiVersion) // api
	admin_route.Admin_route(api)
	//
	app.Listen(":3000")
}
