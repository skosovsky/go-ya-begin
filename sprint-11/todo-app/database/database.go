package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"todo-app/models"
)

type DBInstance struct {
	*gorm.DB
}

var DB DBInstance //nolint:gochecknoglobals,varnamelen // example

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{ //nolint:exhaustruct // example
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	err = db.AutoMigrate(&models.Task{}) //nolint:exhaustruct // example
	if err != nil {
		log.Fatal(err)
	}

	DB = DBInstance{DB: db}
}
