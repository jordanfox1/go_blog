package main

import (
	"go_blog/handlers"
	"go_blog/postgres"
	"go_blog/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	postgres.SetupDBConnection()
	app := fiber.New()

	handler := &handlers.BlogHandlerImpl{}
	routes.SetupBlogRoutes(app, handler)

	app.Listen(":3000")
}
