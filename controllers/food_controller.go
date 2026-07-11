package controller

import "github.com/gofiber/fiber/v3"

func GetFoods(c fiber.Ctx) error   { return c.JSON(fiber.Map{"message": "ok"}) }
func GetFood(c fiber.Ctx) error    { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateFood(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateFood(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
