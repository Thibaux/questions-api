package main

import (
	"github.com/gofiber/fiber/v2"
	"questions-api/database"
	"questions-api/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.getRoutes(app)

	app.Listen(":3000")
}
