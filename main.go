package main

import (
	"fmt"
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

	// Setup Database
	config.SetupDatabase()

	fmt.Println("running on port: " + os.Getenv("APP_PORT"))

	// Run Apps on PORT
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
