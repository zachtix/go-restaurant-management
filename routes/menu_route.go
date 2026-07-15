package route

import (
	controller "restaurant-management/controllers"
	"restaurant-management/middleware"

	"github.com/gofiber/fiber/v3"
)

func MenuRoute(app *fiber.App, h *controller.Controller) {
	menu := app.Group("/menus")
	menu.Use(middleware.Authenticate)
	menu.Get("", middleware.Paginate, h.GetMenus)
	menu.Get("/:menu_id", h.GetMenu)
	menu.Post("", h.CreateMenu)
	menu.Patch("/:menu_id", h.UpdateMenu)
	menu.Delete("/:menu_id", h.DeleteMenu)
}
