package config

import (
	"log"
	"os"
	"testfiber/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectToPostgreSQL() {

	dsn := os.Getenv("PG_DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Perform database migration
	err = db.AutoMigrate(&schema.User{})
	if err != nil {
		log.Fatal(err)
	}
	DB.Db = db
}
