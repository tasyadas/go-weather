package database

import (
	"assignment-3/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "tasyadas"
	dbPort = "5432"
	dbName = "assignment_3"
	db     *gorm.DB
	err    error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, user, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.AutoMigrate(&models.Element{})
}

func GetDB() *gorm.DB {
	return db
}
