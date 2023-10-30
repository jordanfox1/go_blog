package routes

import (
	"go_blog/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupBlogRoutes(app *fiber.App, handler handlers.Handler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/post/all", handler.HandleGetAllPosts)
	app.Get("/post/:id", handler.HandleGetPostByID)

	app.Put("/post/:id", handler.HandleUpdatePostByID)
	app.Delete("/post/:id", handler.HandleDeletePostByID)

	app.Post("/post", handler.HandleCreatePost)
}
