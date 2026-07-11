package route

import (
	"github.com/gofiber/fiber/v3"
	"restaurant-management/controllers"
)

func TableRoute(app *fiber.App) {
	table := app.Group("/tables")
	table.Get("", controller.GetTables)
	table.Get("/:table_id", controller.GetTable)
	table.Post("", controller.CreateTable)
	table.Patch("/:table_id", controller.UpdateTable)
}
