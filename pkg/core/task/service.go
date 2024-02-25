package task

import (
	"errors"
	"strconv"

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

func (svc *Service) GetTaskById() (Task, error) {
	id := svc.Ctx.Params("id")

	idInt, _ := strconv.Atoi(id)

	if idInt == 0 {
		return Task{}, errors.New(_TASK_NOT_FOUND_)
	}

	repo := repository{
		DB:    svc.DB,
		Model: Task{ID: uint(idInt)},
	}

	result, err := repo.GetTaskById()
	if result.ID == 0 {
		err = errors.New(_TASK_NOT_FOUND_)
	}

	return result, err
}

func (svc *Service) DeleteByID(id uint) error {
	repo := repository{
		DB: svc.DB,
	}

	err := repo.Delete(id)

	return err
}

func (svc *Service) UpdateByID(id uint, dataInput TaskAdd) error {
	repo := repository{
		DB:    svc.DB,
		Model: id,
	}

	updateData := Task{
		Name:   dataInput.Name,
		Status: dataInput.Status,
	}

	err := repo.Update(updateData)

	return err
}
