package response

import "github.com/gofiber/fiber/v3"

func Paginated(c fiber.Ctx, data any, page, limit int, total, totalPage int64) error {
	return c.JSON(fiber.Map{
		"message":    "ok",
		"data":       data,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"total_page": totalPage,
	})
}
