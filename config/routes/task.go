package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/pkg/core/task"
)

func Task(taskRoutes fiber.Router) {

	taskRoutes.Get("/", task.GetTasks)

	taskRoutes.Get("/item/:id", func(c *fiber.Ctx) error {
		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Single task",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

	taskRoutes.Post("/", task.AddTask)

	taskRoutes.Delete("/delete-task/:id", func(c *fiber.Ctx) error {
		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Delete task",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

	taskRoutes.Patch("/update-task/:id", func(c *fiber.Ctx) error {
		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Update task",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

}
