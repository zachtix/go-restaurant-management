package route

import (
	controller "restaurant-management/controllers"

	"github.com/gofiber/fiber/v3"
)

func FoodRoute(app *fiber.App, h *controller.Controller) {
	food := app.Group("/foods")
	food.Get("", h.GetFoods)
	food.Get("/:food_id", h.GetFood)
	food.Post("", h.CreateFood)
	food.Patch("/:food_id", h.UpdateFood)
	food.Delete("/:food_id", h.DeleteFood)
}
