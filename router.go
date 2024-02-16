package main

import (
	"testfiber/config"
	handlers "testfiber/handlers/admin"

	"github.com/gofiber/fiber/v2"
)

func EnpRouter(app *fiber.App) {
	api := app.Group(config.Defaults.ApiVersion) // api

	admin := api.Group("/admin")

	admin.Get("/dashboard", handlers.AdminCtrl)

}
