package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/pkg/core/task"
)

func Task(taskRoutes fiber.Router) {

	taskRoutes.Get("/", task.GetTasks)

	taskRoutes.Get("/paginate", task.GetTasksPaginate)

	taskRoutes.Get("/item/:id", task.GetTaskById)

	taskRoutes.Post("/", task.AddTask)

	taskRoutes.Delete("/:id", task.DeleteTask)

	taskRoutes.Patch("/:id", task.UpdateTask)

	taskRoutes.Get("/sample-paginate", task.SamplePaginate)

}
