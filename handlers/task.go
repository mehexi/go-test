package handlers

import (
	db "example/go_api/DB"
	"example/go_api/models"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

var tasks []models.Task

func GetTasks(c fiber.Ctx) error {
	var tasks []models.Task

	if err := db.DB.Find(&tasks).Error; err != nil {
		return ErrorResponse(c, "failed to fetch tasks")
	}

	return SuccessResponse(c, "tasks retrieved", tasks)
}

func AddTask(c fiber.Ctx) error {
	task := new(models.Task)
	if err := c.Bind().Form(task); err != nil {
		fmt.Println("Error parsing form:", err)
		return ErrorResponse(c, "invalid request")
	}

	if err := db.DB.Create(&task).Error; err != nil {
		return ErrorResponse(c, "Failed to create Task")
	}

	return SuccessResponse(c, "task created", task)
}
