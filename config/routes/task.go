package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/pkg/core/task"
)

func Task(taskRoutes fiber.Router) {

	taskRoutes.Get("/", task.GetTasks)

	taskRoutes.Get("/item/:id", task.GetTaskById)

	taskRoutes.Post("/", task.AddTask)

	taskRoutes.Delete("/delete-task/:id", task.DeleteTask)

	taskRoutes.Patch("/update-task/:id", func(c *fiber.Ctx) error {
		// welcome content
		resp := map[string]interface{}{
			"status":  "OK",
			"message": "Update task",
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})

}
