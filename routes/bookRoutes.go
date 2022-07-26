package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huevangxp/go-fiber-rest-api/controllers"
)

func SetRoutesBook(route fiber.Router) {
	route.Get("/", controllers.GetBooks)
	route.Get("/:id", controllers.GetBook)
	route.Post("/", controllers.AddBook)
	route.Put("/:id", controllers.UpdateBook)
	route.Delete("/:id", controllers.DeleteBook)
}
