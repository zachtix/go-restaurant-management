package route

import (
	controller "restaurant-management/controllers"
	"restaurant-management/middleware"

	"github.com/gofiber/fiber/v3"
)

func OrderItemRoute(app *fiber.App, h *controller.Controller) {
	orderItem := app.Group("/order-items")
	orderItem.Use(middleware.Authenticate)
	orderItem.Get("", middleware.Paginate, h.GetOrderItems)
	orderItem.Get("/:order_item_id", h.GetOrderItem)
	orderItem.Get("/order/:order_id", h.GetOrderItemByOrder)
	orderItem.Post("", h.CreateOrderItem)
	orderItem.Patch("/:order_item_id", h.UpdateOrderItem)
	orderItem.Delete("/:order_item_id", h.DeleteOrderItem)
}
