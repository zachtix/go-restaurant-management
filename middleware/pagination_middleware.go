package middleware

import (
	"github.com/gofiber/fiber/v3"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func Paginate(c fiber.Ctx) error {
	page := fiber.Query(c, "page", 1)
	limit := fiber.Query(c, "limit", 5)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	c.Locals("pagination", Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	})

	return c.Next()
}

func GetPagination(c fiber.Ctx) Pagination {
	if p, ok := c.Locals("pagination").(Pagination); ok {
		return p
	}
	return Pagination{Page: 1, Limit: 5, Offset: 0} // fallback เผื่อลืมใส่ middleware
}
