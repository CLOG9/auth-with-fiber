package main

import (
	say_hey "testfiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(say_hey.SayHey)
	EnpRouter(app)

	//
	app.Listen(":3000")
}
