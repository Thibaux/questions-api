package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Thibaux/questions-api/handles"	
)

func getRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
