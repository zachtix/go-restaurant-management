package route

import (
	"github.com/gofiber/fiber/v3"
	"restaurant-management/controllers"
)

func MenuRoute(app *fiber.App) {
	menu := app.Group("/menus")
	menu.Get("", controller.GetMenus)
	menu.Get("/:menu_id", controller.GetMenu)
	menu.Post("", controller.CreateMenu)
	menu.Patch("/:menu_id", controller.UpdateMenu)
}
