package routes

import (
	"go_blog/handlers"
	"go_blog/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupBlogRoutes(app *fiber.App, handler handlers.Handler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/post/all", handler.HandleGetAllPosts)
	app.Get("/post/:id", handler.HandleGetPostByID)

	app.Put("/post/:id", middlewares.ValidatePost(), handler.HandleUpdatePostByID)
	app.Delete("/post/:id", handler.HandleDeletePostByID)

	app.Post("/post", middlewares.ValidatePost(), handler.HandleCreatePost)
}
