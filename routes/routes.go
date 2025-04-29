package routes

import (
	"goapp/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App) {
	api := app.Group("/api/v1")

	author := api.Group("/author")
	author.Post("/", handlers.CreateAuthor)
	author.Put("/:id", handlers.UpdateAuthor)
	author.Get("/", handlers.ListAuthors)
	author.Get("/:id", handlers.GetAuthor)
	author.Delete("/:id", handlers.DeleteAuthor)

	//book := api.Group("/book")
	//book.Post("/", handlers.CreateBook)
	//book.Put("/:id", handlers.UpdateBook)
	//book.Get("/", handlers.ListBooks)
	//book.Get("/:id", handlers.GetBook)
	//book.Delete("/:id", handlers.DeleteBook)
}