package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/handlers"
)

func getRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
