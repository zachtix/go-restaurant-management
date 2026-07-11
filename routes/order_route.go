package route

import (
	"github.com/gofiber/fiber/v3"
	"restaurant-management/controllers"
)

func OrderRoute(app *fiber.App) {
	order := app.Group("/orders")
	order.Get("", controller.GetOrders)
	order.Get("/:order_id", controller.GetOrder)
	order.Post("", controller.CreateOrder)
	order.Patch("/:order_id", controller.UpdateOrder)
}
