package main

import (
	"log"
	"os"

	controller "restaurant-management/controllers"
	"restaurant-management/database"
	route "restaurant-management/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	app := SetupFiber()
	app.Listen(":" + PORT)
}

func SetupFiber() *fiber.App {
	db := database.GormInitialize()

	h := &controller.Controller{DB: db}

	app := fiber.New()

	route.UserRoute(app, h)
	route.FoodRoute(app, h)
	route.OrderRoute(app, h)
	route.OrderItemRoute(app, h)
	route.MenuRoute(app, h)
	route.TableRoute(app, h)
	route.InvoiceRoute(app, h)

	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found"})
	})

	return app
}
