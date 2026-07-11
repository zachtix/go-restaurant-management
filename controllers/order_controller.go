package controller

import "github.com/gofiber/fiber/v3"

func GetOrders(c fiber.Ctx) error   { return c.JSON(fiber.Map{"message": "ok"}) }
func GetOrder(c fiber.Ctx) error    { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateOrder(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateOrder(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
