package db

import (
	"fmt"
	"log"

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

	return db
}
