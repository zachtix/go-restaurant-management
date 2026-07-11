package main

import (
	"fmt"
	"log"
	"os"

	// "restaurant-management/controller"
	"restaurant-management/database"
	// "restaurant-management/helper"
	// "restaurant-management/middleware"
	// "restaurant-management/models"
	"restaurant-management/routes"

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

	fmt.Print(db)

	app := fiber.New()

	route.UserRoute(app)
	route.FoodRoute(app)
	route.OrderRoute(app)
	route.OrderItemRoute(app)
	route.MenuRoute(app)
	route.TableRoute(app)
	route.InvoiceRoute(app)

	return app
}
