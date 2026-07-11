package main

import "github.com/gofiber/fiber/v3"

func main() {

}

func SetupFiber() *fiber.App {
	app := fiber.New()

	app.Listen(":8000")

	return app
}
