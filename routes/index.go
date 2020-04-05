package routes

import "github.com/gofiber/fiber"

// Index route
func Index(c *fiber.Ctx) {
	_ = c.Render("index", fiber.Map{
		"socket_address": "ws://localhost:3000/ws",
	})
}

