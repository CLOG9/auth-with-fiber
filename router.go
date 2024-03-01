package main

import (
	"testfiber/config"
	handlers "testfiber/handlers/auth"
	"testfiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func EnpRouter(app *fiber.App) {
	api := app.Group(config.Defaults.ApiVersion) // api

	auth := api.Group("/auth")

	auth.Post("/login", handlers.LoginCtrl)
	auth.Post("/register", handlers.RegisterCtrl)
	auth.Get("/logout", middlewares.Authenticate, handlers.LogoutCtrl)

	auth.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hey")
	})
}
