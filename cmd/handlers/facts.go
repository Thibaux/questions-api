package handlers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
    return c.SendString("Div Rhino Trivia AppDiv Rhino Trivia AppDiv Rhino Trivia AppDiv Rhino Trivia AppDiv Rhino Trivia Apps!")
}
