package main

import (
	"testfiber/config"
	say_hey "testfiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectToPostgreSQL()
	app.Use(say_hey.SayHey)
	EnpRouter(app)

	//
	app.Listen(":3000")
}
