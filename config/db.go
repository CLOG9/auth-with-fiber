package config

import (
	"fmt"
	"log"
	"testfiber/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectToPostgreSQL() {
	dsn := "user=root password=roievn2zv7rz6hr9th3q4g dbname=pgs_db host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("im heeeerrrrr")
		log.Fatal(err)
	}
	// Perform database migration
	err = db.AutoMigrate(&schema.User{})
	fmt.Println("done")
	if err != nil {
		log.Fatal(err)
	}
	DB.Db = db
}
