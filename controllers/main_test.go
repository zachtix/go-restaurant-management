package controller

import (
	model "restaurant-management/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}
	if err := db.AutoMigrate(
		&model.User{},
		&model.Order{},
		&model.OrderItem{},
		&model.Food{},
		&model.Invoice{},
		&model.Menu{},
		&model.Note{},
		&model.Table{},
	); err != nil {
		t.Fatalf("Migrate: %v", err)
	}
	return db
}

func setupAppTest(db *gorm.DB) (*fiber.App, *Controller) {
	h := &Controller{DB: db}
	app := fiber.New()
	return app, h
}
