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
