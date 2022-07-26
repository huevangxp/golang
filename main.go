package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huevangxp/go-fiber-rest-api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "github.com/huevangxp/go-fiber-rest-api",
		})
	})

	api := app.Group("/api")

	routes.SetRoutesBook(api.Group("/books"))

}

func main() {

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":7000")
}
