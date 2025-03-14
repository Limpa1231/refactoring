package database

import (
	"firstRest/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres dbname=fark sslmode=disable password=root"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Автомиграция схемы
	err = DB.AutoMigrate(&models.Message{}, &models.User{}) // Добавьте &models.User{}
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
