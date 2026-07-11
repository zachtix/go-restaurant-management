package controller

import "github.com/gofiber/fiber/v3"

func GetInvoices(c fiber.Ctx) error   { return c.JSON(fiber.Map{"message": "ok"}) }
func GetInvoice(c fiber.Ctx) error    { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateInvoice(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func UpdateInvoice(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
