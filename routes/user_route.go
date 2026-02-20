package routes

import (
	"example/go_api/handlers"
	"example/go_api/middleware"

	"github.com/gofiber/fiber/v3"
)

func UserRoutes(app fiber.App) {
	api := app.Group("/api")
	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.Register)

	private := app.Group("/api", middleware.AuthRequired)
	private.Get("/me", handlers.GetUserInfo)
	private.Patch("/update-password", handlers.UpdatePassword)
}
