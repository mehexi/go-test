package handlers

import (
	db "example/go_api/DB"
	"example/go_api/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var tasks []models.Task

func GetTasks(c fiber.Ctx) error {
	var tasks []models.Task
	userID := c.Locals("user_id").(uint)

	if err := db.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
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

	userID := c.Locals("user_id").(uint)
	task.UserID = userID

	if err := db.DB.Create(&task).Error; err != nil {
		return ErrorResponse(c, "Failed to create Task")
	}

	return SuccessResponse(c, "task created", task)
}

func DeleteTask(c fiber.Ctx) error {
	tasks := new(models.Task)
	idStr := c.Params("id")
	userID := c.Locals("user_id").(uint)

	if idStr == "" {
		return ErrorResponse(c, "No Query is provided ")
	}

	result := db.DB.Where("id = ? AND user_id = ?", idStr, userID).Delete(tasks, idStr)

	if result.Error != nil {
		return ErrorResponse(c, "Failed to delete Task")
	}

	if result.RowsAffected == 0 {
		return ErrorResponse(c, "Task not found")
	}

	return SuccessResponse(c, "Task Has been deleted", nil)
}

func EditTask(c fiber.Ctx) error {
	idStr := c.Params("id")
	userID := c.Locals("user_id").(uint)
	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		return ErrorResponse(c, "invalid id")
	}

	var task models.Task

	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&task, id).Error; err != nil {
		return ErrorResponse(c, "Task not found")
	}

	if err := c.Bind().Form(&task); err != nil {
		return ErrorResponse(c, "invalid reqest body")
	}

	if err := db.DB.Where("id = ? AND user_id= ?", id, userID).Save(&task).Error; err != nil {
		return ErrorResponse(c, "Couldnt update the task")
	}

	return SuccessResponse(c, "task updated", task)
}
