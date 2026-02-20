package handlers

import (
	db "example/go_api/DB"
	"example/go_api/models"

	"github.com/gofiber/fiber/v3"
)

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func GetUserInfo(c fiber.Ctx) error {
	var user models.User
	id, ok := c.Locals("user_id").(uint)

	if !ok {
		return ErrorResponse(c, "unauthorized")
	}

	if err := db.DB.Preload("Task").Where("id = ?", id).First(&user).Error; err != nil {
		return ErrorResponse(c, "user not found")
	}

	return c.JSON(user)
}

func UpdatePassword(c fiber.Ctx) error {
	var input UpdatePasswordInput
	var user models.User

	id := c.Locals("user_id").(uint)

	if err := c.Bind().Body(&input); err != nil {
		return ErrorResponse(c, "Invalid request")
	}

	if input.OldPassword == "" || input.NewPassword == "" {
		return ErrorResponse(c, "Old and new password are required")
	}

	if err := db.DB.First(&user, id).Error; err != nil {
		return ErrorResponse(c, "User not found")
	}

	if err := user.UpdatePassword(input.OldPassword, input.NewPassword); err != nil {
		return ErrorResponse(c, err.Error())
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return ErrorResponse(c, "Failed to update password")
	}

	return SuccessResponse(c, "Password updated successfully", nil)
}
