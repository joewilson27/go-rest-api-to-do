package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/config"
)

func main() {
	// Setup App
	app := fiber.New()

	// Seup Routes
	config.SetupRoutes(app)

	// check and load ENV
	config.LoadEnvir()

	// Run Apps on PORT
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
