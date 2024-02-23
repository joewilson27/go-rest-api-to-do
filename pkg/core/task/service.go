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

	dataTask := Task{Name: data.Name, Status: data.Status}

	repo := repository{
		DB:    svc.DB,
		Model: &dataTask,
	}

	err := repo.Create()
	return err
}

func CreateTask(name, status string) (Task, error) {
	var newTask = Task{Name: name, Status: status}

	err := db.PG.Create(&Task{Name: name, Status: status})
	if err != nil {
		return newTask, err.Error
	}

	return newTask, nil
}
