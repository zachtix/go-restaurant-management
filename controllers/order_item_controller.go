package controller

import "github.com/gofiber/fiber/v3"

func GetOrderItems(c fiber.Ctx) error       { return c.JSON(fiber.Map{"message": "ok"}) }
func GetOrderItem(c fiber.Ctx) error        { return c.JSON(fiber.Map{"message": "ok"}) }
func GetOrderItemByOrder(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateOrderItem(c fiber.Ctx) error     { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateOrderItem(c fiber.Ctx) error     { return c.JSON(fiber.Map{"message": "ok"}) }
