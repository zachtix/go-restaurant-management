package route

import (
	controller "restaurant-management/controllers"
	"restaurant-management/middleware"

	"github.com/gofiber/fiber/v3"
)

func OrderRoute(app *fiber.App, h *controller.Controller) {
	order := app.Group("/orders")
	order.Use(middleware.Authenticate)
	order.Get("", middleware.Paginate, h.GetOrders)
	order.Get("/:order_id", h.GetOrder)
	order.Post("", h.CreateOrder)
	order.Patch("/:order_id", h.UpdateOrder)
	order.Delete("/:order_id", h.DeleteOrder)
}
