package main

import (
	"gogarden/bookmark"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/bookmark", bookmark.GetAllBookmarks)
	app.Post("/api/bookmark", bookmark.SaveBookmark)
}

func main() {
	// Create a new instance of Fiber
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":3000")
}
