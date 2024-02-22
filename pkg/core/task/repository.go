package task

import (
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB    *gorm.DB
	Model any
}

func (repo *repository) Create() error {
	data := repo.Model
	fmt.Println("print before ", data)
	err := repo.DB.Create(&data)
	return err.Error
}
