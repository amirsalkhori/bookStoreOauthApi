package authDB

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	username := "unicorn_user"
	password := "magical_password"
	host := "127.0.0.1"
	port := "5432"
	db := "rainbow_database"

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host, port, db, username, password)

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	sqlDB, err := client.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established successfully!")

	return client, nil
}
