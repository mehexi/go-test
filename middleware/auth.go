package middleware

import (
	"example/go_api/utils"
	"github.com/gofiber/fiber/v3"
	"strings"
)

func AuthRequired(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized - No token provided",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized - Invalid token format",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized - Invalid or expired token",
		})
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("username", claims.Username)
	c.Locals("email", claims.Email)

	return c.Next()
}
