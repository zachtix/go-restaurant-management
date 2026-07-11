package route

import (
	"restaurant-management/controllers"
	"restaurant-management/middleware"

	"github.com/gofiber/fiber/v3"
)

func UserRoute(app *fiber.App, h *controller.Controller) {
	auth := app.Group("/auth")
	user := app.Group("/users")
	user.Use(middleware.Authenticate)
	user.Get("/", h.GetUsers)
	user.Get("/:user_id", h.GetUser)
	auth.Post("/signup", h.CreateUser)
	auth.Post("/signin", h.LoginUser)
}
