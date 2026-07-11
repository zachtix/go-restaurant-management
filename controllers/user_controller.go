package controller

import (
	"github.com/gofiber/fiber/v3"
)

func GetUsers(c fiber.Ctx) error   { return c.JSON(fiber.Map{"message": "ok"}) }
func GetUser(c fiber.Ctx) error    { return c.JSON(fiber.Map{"message": "ok"}) }
func CreateUser(c fiber.Ctx) error { return c.JSON(fiber.Map{"message": "ok"}) }
func LoginUser(c fiber.Ctx) error  { return c.JSON(fiber.Map{"message": "ok"}) }
