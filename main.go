package main

import (
	db "example/go_api/DB"
	"example/go_api/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	db.Connect()
	db.Migrate()
	app := fiber.New()

	routes.UserRoutes(*app)
	routes.TaskRoutes(*app)

	log.Fatal(app.Listen(":3000"))

}
