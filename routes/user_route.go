package route

import (
	"restaurant-management/controllers"

	"github.com/gofiber/fiber/v3"
)

func UserRoute(app *fiber.App) {
	auth := app.Group("/auth")
	user := app.Group("/users")
	user.Get("/", controller.GetUsers)
	user.Get("/:user_id", controller.GetUser)
	auth.Post("/signup", controller.CreateUser)
	auth.Post("/signin", controller.LoginUser)
}
