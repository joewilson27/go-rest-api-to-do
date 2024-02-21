package config

import (
	"fmt"

	"github.com/joewilson27/config/db"
	"github.com/joewilson27/pkg/core/task"
)

func SetupDatabase() {
	if err := db.InitDB(); err == nil {

		// Create the schema
		db.PG.Exec("CREATE SCHEMA IF NOT EXISTS todo")

		db.PG.AutoMigrate(
			&task.Task{},
		)
		fmt.Println("Migrate database successfully")
	}
}
