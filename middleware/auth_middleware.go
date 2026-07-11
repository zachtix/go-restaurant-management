package middleware

import (
	"os"
	"restaurant-management/helper"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Authenticate(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing authorization header"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid authorization format"})
	}

	claims, err := helper.VerifyJWT(tokenString, os.Getenv("SECRET_ACCESS"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
	}

	c.Locals("user", claims)
	return c.Next()
}
