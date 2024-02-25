package task

import (
	"gorm.io/gorm"
)

type repository struct {
	DB    *gorm.DB
	Model any
}

func (repo *repository) Create() error {
	err := repo.DB.Create(repo.Model)
	return err.Error
}

func (repo *repository) GetTasks() ([]*Task, error) {
	tasks := []*Task{}

	err := repo.DB.Find(&tasks)

	if err.Error != nil {
		return nil, err.Error
	}

	return tasks, nil
}

func (repo *repository) GetTaskById() (Task, error) {
	var task Task

	result := repo.DB.Where(repo.Model).Limit(1).Find(&task)

	return task, result.Error
}

func (repo *repository) Delete(id uint) error {
	var task Task
	//result := repo.DB.Delete(&Task{ID: id})
	result := repo.DB.Where("id = ?", id).Delete(&task)
	return result.Error
}

func (repo *repository) Update(data Task) error {

	result := repo.DB.Where("id = ?", repo.Model).Updates(data)

	return result.Error
}
