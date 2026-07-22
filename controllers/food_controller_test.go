package controller

import (
	"encoding/json"
	"net/http/httptest"
	"restaurant-management/middleware"
	model "restaurant-management/models"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFood(t *testing.T) {
	db := setupTestDB(t)
	app, h := setupAppTest(db)
	app.Get("/foods/:food_id", h.GetFood)

	// Seed data
	db.Create(&model.Food{Name: "Pad Thai", Price: 60, Food_image: "img.png", Menu_id: 1})

	t.Run("Found Data", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/foods/1", nil)
		resq, err := app.Test(req)

		require.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resq.StatusCode)
	})

	t.Run("Not Found -> 404", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/foods/2", nil)
		resq, err := app.Test(req)

		require.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resq.StatusCode)
	})

	t.Run("ID invalid -> 400", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/foods/ABC", nil)
		resq, err := app.Test(req)

		require.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resq.StatusCode)
	})
}

func TestCreateFood(t *testing.T) {
	db := setupTestDB(t)
	app, h := setupAppTest(db)
	app.Post("/foods", h.CreateFood)

	// Seed data
	db.Create(&model.Menu{Name: "Main", Category: "Lunch"})

	t.Run("Create success -> 201", func(t *testing.T) {
		body := `{"name":"Sum Tum","price":50,"food_image":"s.png","menu_id":1}`
		req := httptest.NewRequest("POST", "/foods", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resq, err := app.Test(req)

		require.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resq.StatusCode)

		var count int64
		db.Model(&model.Food{}).Count(&count)
		assert.Equal(t, int64(1), count)
	})

	t.Run("Not Found Menu ID -> 404", func(t *testing.T) {
		body := `{"name":"Sum Tum","price":50,"food_image":"s.png","menu_id":2}`
		req := httptest.NewRequest("POST", "/foods", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resq, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resq.StatusCode)
	})

	t.Run("Validation Fail -> 400", func(t *testing.T) {
		body := `{"price":50,"food_image":"s.png","menu_id":1}`
		req := httptest.NewRequest("POST", "/foods", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resq, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resq.StatusCode)
	})
}

func TestGetFoodsPagination(t *testing.T) {
	db := setupTestDB(t)
	h := &Controller{DB: db}
	app := fiber.New()
	app.Get("/foods", middleware.Paginate, h.GetFoods)

	for i := 0; i < 7; i++ {
		db.Create(&model.Food{Name: "F", Price: 10, Food_image: "i", Menu_id: 1})
	}

	req := httptest.NewRequest("GET", "/foods?page=1&limit=5", nil)
	resq, err := app.Test(req)
	require.NoError(t, err)

	var out struct {
		Total     int64 `json:"total"`
		TotalPage int64 `json:"total_page"`
	}

	require.NoError(t, json.NewDecoder(resq.Body).Decode(&out))
	assert.Equal(t, int64(7), out.Total)
	assert.Equal(t, int64(2), out.TotalPage)
}
