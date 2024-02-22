package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PG *gorm.DB

func InitDB() error {

	dsn := "host=" + os.Getenv("PG_HOST") + " user=" + os.Getenv("PG_USER") + " password=" + os.Getenv("PG_PASSWORD") +
		" dbname=" + os.Getenv("PG_DB_NAME") + " port=" + os.Getenv("PG_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "todo.",
			SingularTable: false,
		},
	})

	if err != nil {
		log.Println("Can't connect to database because : ", err.Error())
		return err
	}

	PG = database
	fmt.Println("Success Connect to DB")

	return nil
}
