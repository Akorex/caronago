package database

import (
	"log"

	"github.com/Akorex/caronago/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func Connect(dsn string){
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database successfully")

	DB = db

	db.AutoMigrate(&models.User{})
}