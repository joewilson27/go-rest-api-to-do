package task

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/common/model"
	"github.com/joewilson27/common/response"
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

	resp, errAdd := svc.Add(newTask)
	if errAdd != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.DataAddedFailed + ": " + errAdd.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusAccepted).JSON(model.AppResponse{
		Data: resp,
		Meta: model.Meta{Message: response.DataAddedSuccessfully},
	})
}

func GetTasks(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTasks()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.GetDataFailed + " : " + err.Error()},
		}
		return c.Status(fiber.StatusBadRequest).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.GetDataSuccessfully},
	})
}

func GetTaskById(c *fiber.Ctx) error {

	svc := Service{
		DB:  db.PG,
		Ctx: c,
	}

	result, err := svc.GetTaskById()
	if err != nil {
		dataReturnError := model.AppResponse{
			Data: nil,
			Meta: model.Meta{Message: response.DataNotFound + " " + err.Error()},
		}

		return c.Status(fiber.StatusNotFound).JSON(dataReturnError)
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{
		Data: result,
		Meta: model.Meta{Message: response.StatusSuccess},
	})
}
