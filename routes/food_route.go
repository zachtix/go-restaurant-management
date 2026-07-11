package route

import (
	"github.com/gofiber/fiber/v3"
	"restaurant-management/controllers"
)

func FoodRoute(app *fiber.App) {
	food := app.Group("/foods")
	food.Get("", controller.GetFoods)
	food.Get("/:food_id", controller.GetFood)
	food.Post("", controller.CreateFood)
	food.Patch("/:food_id", controller.UpdateFood)
}
