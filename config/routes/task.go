package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/pkg/core/task"
)

func Task(taskRoutes fiber.Router) {

	taskRoutes.Get("/", func(c *fiber.Ctx) error {
		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "All tasks",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

	taskRoutes.Get("/item/:id", func(c *fiber.Ctx) error {
		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Single task",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

	taskRoutes.Post("/", task.AddTask)
	// taskRoutes.Post("/add-task", func(c *fiber.Ctx) error {
	// 	// welcome content
	// 	resp := map[string]interface{}{
	// 		"status":  "OK",
	// 		"message": "Add task",
	// 	}

	// 	return c.Status(fiber.StatusOK).JSON(resp)
	// })

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
