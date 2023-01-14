package main

import (
	"github.com/gofiber/fiber/v2"
    "github.com/Thibaux/questions-api/cmd/database"
    "github.com/Thibaux/questions-api/cmd/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

    routes.GetRoutes(app)

	app.Listen(":3000")
}
