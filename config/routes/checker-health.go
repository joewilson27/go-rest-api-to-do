package routes

import "github.com/gofiber/fiber/v2"

func CheckerHealth(checkerHealthRoutes fiber.Router) {
	checkerHealthRoutes.Get("/", func(c *fiber.Ctx) error {

		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Welcome to the Todo-List-API",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})
}
