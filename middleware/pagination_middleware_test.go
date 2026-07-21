package middleware

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestPaginate(t *testing.T) {
	app := fiber.New()
	app.Get("/", Paginate, func(c fiber.Ctx) error {
		p := GetPagination(c)
		return c.JSON(p)
	})

	req := httptest.NewRequest("GET", "/?page=2&limit=500", nil)
	resq, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resq.StatusCode != 200 {
		t.Fatalf("Status = %d", resq.StatusCode)
	}

	var p Pagination
	err = json.NewDecoder(resq.Body).Decode(&p)
	assert.NoError(t, err)

	assert.Equal(t, 2, p.Page)
	assert.Equal(t, 100, p.Limit)
	assert.Equal(t, 100, p.Offset)
}
