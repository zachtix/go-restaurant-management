package route

import (
	controller "restaurant-management/controllers"
	"restaurant-management/middleware"

	"github.com/gofiber/fiber/v3"
)

func TableRoute(app *fiber.App, h *controller.Controller) {
	table := app.Group("/tables")
	table.Use(middleware.Authenticate)
	table.Get("", middleware.Paginate, h.GetTables)
	table.Get("/:table_id", h.GetTable)
	table.Post("", h.CreateTable)
	table.Patch("/:table_id", h.UpdateTable)
	table.Delete("/:table_id", h.DeleteTable)
}
