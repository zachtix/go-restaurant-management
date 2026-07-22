package controller

import (
	"io"
	"net/http/httptest"
	model "restaurant-management/models"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInvoice(t *testing.T) {
	db := setupTestDB(t)
	app, h := setupAppTest(db)
	app.Get("/invoices/:invoice_id", h.GetInvoice)

	// Seed data
	db.Create(&model.Table{Number_of_guests: 2, Table_Number: 1})
	db.Create(&model.Order{Order_date: time.Now(), Table_id: 1})
	db.Create(&model.Invoice{Order_id: 1, Payment_method: "CARD", Payment_status: "PENDING", Payment_due_date: time.Now().Add(time.Hour + 1)})

	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{"Get invoice -> 200", "/invoices/1", fiber.StatusOK},
		{"Not found invoice -> 404", "/invoices/2", fiber.StatusNotFound},
		{"Invalid ID -> 400", "/invoices/ABC", fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			resq, err := app.Test(req)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, resq.StatusCode)
		})
	}
}

func TestCreateInvoice(t *testing.T) {
	db := setupTestDB(t)
	app, h := setupAppTest(db)
	app.Post("/invoices", h.CreateInvoice)

	// Seed data
	db.Create(&model.Table{Number_of_guests: 2, Table_Number: 1})
	db.Create(&model.Order{Order_date: time.Now(), Table_id: 1})

	tests := []struct {
		name     string
		body     string
		expected int
	}{
		{"Create success -> 201", `{"order_id":1,"payment_method":"CARD"}`, fiber.StatusCreated},
		{"Validation fail -> 400", `{payment_method":"CARD"}`, fiber.StatusBadRequest},
		{"Not found order id -> 400", `{"order_id":999,"payment_method":"CARD"}`, fiber.StatusBadRequest},
		{"Order id duplicate -> 409", `{"order_id":1,"payment_method":"CARD"}`, fiber.StatusConflict},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/invoices", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")

			resq, err := app.Test(req)

			if tt.expected != resq.StatusCode {
				body, _ := io.ReadAll(resq.Body)
				t.Logf("[%s] status=%d body=%s", tt.name, resq.StatusCode, string(body))
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, resq.StatusCode)
		})
	}
}
