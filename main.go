package main

import (
	"log"
	"testfiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	config.GoogleConfig()
	// db
	config.ConnectToPostgreSQL()

	// session
	config.SessCnf()

	// router
	EnpRouter(app)

	//
	app.Listen(":3000")
}
