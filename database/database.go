package database

import (
	"fmt"
	"os"
	"restaurant-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormInitialize() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&model.User{},
		&model.Order{},
		&model.OrderItem{},
		&model.Food{},
		&model.Invoice{},
		&model.Menu{},
		&model.Note{},
		&model.Table{},
	)

	return db
}
