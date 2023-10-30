package main

import (
	"go_blog/handlers"
	"go_blog/routes"
	"go_blog/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := storage.SetupDBConnection() // uses postgres as a db
	app := fiber.New()

	postgresStorage := &storage.PostgresStorage{DB: db}

	// pass in the postgres storage implementation to the handler
	handler := &handlers.BlogHandlerImpl{
		Storage: postgresStorage,
	}

	routes.SetupBlogRoutes(app, handler)

	app.Listen(":3000")
}
