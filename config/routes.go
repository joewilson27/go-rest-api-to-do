package config

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// api
	api := app.Group("/api")

	api.Get("/checker-health", func(c *fiber.Ctx) error {

		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Welcome to the Todo-List-API",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

}
