package main

import (
	"github.com/gofiber/fiber/v2"
)

func getRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
