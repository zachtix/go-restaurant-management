package controller

import "github.com/gofiber/fiber/v3"

func GetMenus(c fiber.Ctx) error   { return c.JSON(fiber.Map{"message": "ok"}) }
func GetMenu(c fiber.Ctx) error    { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateMenu(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateMenu(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
