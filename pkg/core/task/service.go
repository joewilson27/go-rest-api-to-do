package task

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/config/db"
	"gorm.io/gorm"
)

type Service struct {
	DB  *gorm.DB
	Ctx *fiber.Ctx
}

// Add
func (svc *Service) Add(data TaskAdd) error {

	dataTask := Task{
		Name:   data.Name,
		Status: data.Status,
	}

	repo := repository{
		DB:    db.PG,
		Model: dataTask,
	}

	err := repo.Create()
	return err
}
