package routes

import (
	"github.com/gofiber/fiber/v2"

    "github.com/Thibaux/questions-api/cmd/handlers"
)

func GetRoutes(app *fiber.App) {
    app.Get("/", handlers.Index)
}
