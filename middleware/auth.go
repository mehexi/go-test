package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func AuthRequired(c fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized - No token provided",
		})
	}

	if token != "Bearer secret-token" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized - Invalid token",
		})
	}

	return c.Next()
}
