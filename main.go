package main

import (
	"go_blog/postgres"
	"go_blog/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	postgres.SetupDBConnection()
	app := fiber.New()
	routes.SetupBlogRoutes(app)

	app.Listen(":3000")
}
