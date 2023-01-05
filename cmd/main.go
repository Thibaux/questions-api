package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Thibaux/questions-api/database"
	"github.com/Thibaux/questions-api/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.getRoutes(app)

	app.Listen(":3000")
}
