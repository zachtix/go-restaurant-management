package controller

import "github.com/gofiber/fiber/v3"

func GetTables(c fiber.Ctx) error   { return c.JSON(fiber.Map{"message": "ok"}) }
func GetTable(c fiber.Ctx) error    { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateTable(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateTable(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
