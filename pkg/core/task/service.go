package task

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service struct {
	DB  *gorm.DB
	Ctx *fiber.Ctx
}

// Add
func (svc *Service) Add(data TaskAdd) (Task, error) {

	dataTask := Task{Name: data.Name, Status: data.Status}

	repo := repository{
		DB:    svc.DB,
		Model: &dataTask,
	}

	err := repo.Create()
	return dataTask, err
}

func (svc *Service) GetTasks() ([]*Task, error) {

	repo := repository{
		DB: svc.DB,
	}

	result, err := repo.GetTasks()

	if err != nil {
		return nil, err
	}

	return result, err
}
