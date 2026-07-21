package middleware

import (
	"net/http/httptest"
	"os"
	"restaurant-management/helper"
	model "restaurant-management/models"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func setupAuthApp() *fiber.App {
	app := fiber.New()
	app.Get("/protected", Authenticate, func(c fiber.Ctx) error {
		return c.SendString("ok")
	})
	return app
}

func TestAuthenticate(t *testing.T) {
	secret := "jwt-secret"
	secret = os.Getenv("SECRET_ACCESS")

	validToken, _ := helper.GenerateJWT(model.User{Email: "test@mail.com", First_name: "John"}, secret, 1)

	tests := []struct {
		name       string
		authHeader string
		wantStatus int
	}{
		{"No header", "", fiber.StatusUnauthorized},
		{"Token invalid", "Bearer garbage-token", fiber.StatusUnauthorized},
		{"Token valid", "Bearer " + validToken, fiber.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := setupAuthApp()
			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status = %d; want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}
