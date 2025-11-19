package db

import (
	"fmt"
	"log"

	"github.com/slava-abramoff/question-answer-api/application/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=app password=app dbname=app port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get sql.DB from gorm:", err)
		return nil
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Database is not reachable:", err)
		return nil
	}

	fmt.Println("Database connected successfully")

	// Автомиграция моделей
	if err := db.AutoMigrate(&models.Question{}, &models.Answer{}); err != nil {
		log.Println("AutoMigrate failed:", err)
		return nil
	}

	return db
}
