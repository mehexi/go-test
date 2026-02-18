package handlers

import (
	db "example/go_api/DB"
	"example/go_api/models"
	"example/go_api/utils"

	"github.com/gofiber/fiber/v3"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register new user
func Register(c fiber.Ctx) error {
	req := new(RegisterRequest)

	if err := c.Bind().JSON(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user already exists
	var existingUser models.User
	if err := db.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "User with this email already exists",
		})
	}

	// Create new user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
	}

	// Hash password
	if err := user.HashPassword(req.Password); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Save to db
	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Generate token
	token, err := utils.GeneartedTokane(user.ID, user.Username, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User registered successfully",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
			"token": token,
		},
	})
}

// Login user
func Login(c fiber.Ctx) error {
	req := new(LoginRequest)

	if err := c.Bind().JSON(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find user by email
	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Check password
	if err := user.CheckPassword(req.Password); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate token
	token, err := utils.GeneartedTokane(user.ID, user.Username, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
			"token": token,
		},
	})
}

// Get current user info
func GetMe(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	username := c.Locals("username").(string)
	email := c.Locals("email").(string)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":       userID,
			"username": username,
			"email":    email,
		},
	})
}
