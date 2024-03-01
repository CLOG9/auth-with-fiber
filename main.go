package main

import (
	"testfiber/config"
	"testfiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// db
	config.ConnectToPostgreSQL()

	// session
	config.SessCnf()

	app.Use(
		[]string{"/api/v1/auth/logout", "/api/v1/auth/"},
		middlewares.Authenticate,
	)
	// router
	EnpRouter(app)

	//
	app.Listen(":3000")
}
