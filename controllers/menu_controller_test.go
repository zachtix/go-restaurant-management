package controller

import (
	"net/http/httptest"
	model "restaurant-management/models"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetMenu(t *testing.T) {
	db := setupTestDB(t)
	app, h := setupAppTest(db)
	app.Get("/menus/:menu_id", h.GetMenu)

	// Seed data
	db.Create(&model.Menu{Name: "Soup", Category: "Soup"})

	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{"Get menu -> 200", "/menus/1", fiber.StatusOK},
		{"Not Found -> 404", "/menus/2", fiber.StatusNotFound},
		{"Invalid Menu ID -> 400", "/menus/ABC", fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			resq, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, resq.StatusCode)
		})
	}
}
