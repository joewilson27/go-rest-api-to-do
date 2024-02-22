package task

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/common/model"
	"github.com/joewilson27/config/db"
)

func AddTask(c *fiber.Ctx) error {
	var newTask TaskAdd

	if err := c.BodyParser(&newTask); err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	svc := Service{
		DB: db.PG,
	}

	if errAdd := svc.Add(newTask); errAdd != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: "Error add task. " + errAdd.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusAccepted).JSON(model.AppResponse{
		Data: nil,
		Meta: model.Meta{Message: newTask.Name},
	})
}
