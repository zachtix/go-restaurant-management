package route

import (
	"github.com/gofiber/fiber/v3"
	"restaurant-management/controllers"
)

func OrderItemRoute(app *fiber.App) {
	orderItem := app.Group("/order-items")
	orderItem.Get("", controller.GetOrderItems)
	orderItem.Get("/:order_item_id", controller.GetOrderItem)
	orderItem.Get("/order/:order_id", controller.GetOrderItemByOrder)
	orderItem.Post("", controller.CreateOrderItem)
	orderItem.Patch("/:order_item_id", controller.UpdateOrderItem)
}
