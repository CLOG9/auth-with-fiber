package main

import (
	"testfiber/config"
	handlers "testfiber/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func EnpRouter(app *fiber.App) {
	api := app.Group(config.Defaults.ApiVersion) // api

	auth := api.Group("/auth")

	auth.Post("/login", handlers.LoginCtrl)
	auth.Post("/register", handlers.RegisterCtrl)

}
