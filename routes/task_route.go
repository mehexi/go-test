package routes

import (
	"example/go_api/handlers"
	"example/go_api/middleware"

	"github.com/gofiber/fiber/v3"
)

func TaskRoutes(app fiber.App) {
	api := app.Group("/api")

	api.Get("/", handlers.Root)

	api.Use("/task", middleware.AuthRequired)
	api.Get("/task", handlers.GetTasks)
	api.Post("/task", handlers.AddTask)
	api.Delete("/task/:id", handlers.DeleteTask)
	api.Patch("/task/:id", handlers.EditTask)
}
