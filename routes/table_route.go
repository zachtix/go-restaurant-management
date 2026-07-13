package route

import (
	controller "restaurant-management/controllers"

	"github.com/gofiber/fiber/v3"
)

func TableRoute(app *fiber.App, h *controller.Controller) {
	table := app.Group("/tables")
	table.Get("", h.GetTables)
	table.Get("/:table_id", h.GetTable)
	table.Post("", h.CreateTable)
	table.Patch("/:table_id", h.UpdateTable)
	table.Delete("/:table_id", h.DeleteTable)
}
