package route

import (
	controller "restaurant-management/controllers"

	"github.com/gofiber/fiber/v3"
)

func OrderRoute(app *fiber.App, h *controller.Controller) {
	order := app.Group("/orders")
	order.Get("", h.GetOrders)
	order.Get("/:order_id", h.GetOrder)
	order.Post("", h.CreateOrder)
	order.Patch("/:order_id", h.UpdateOrder)
	order.Delete("/:order_id", h.DeleteOrder)
}
