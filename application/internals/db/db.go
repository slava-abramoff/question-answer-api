package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

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
