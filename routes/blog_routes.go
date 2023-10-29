package routes

import "github.com/gofiber/fiber/v2"

func SetupBlogRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/post/all", getAllPosts)
	app.Get("/post/:id", getPostByID)

	app.Put("/post/:id", updatePostByID)
	app.Delete("/post/:id", deletePostByID)

	app.Post("/post", createPost)
}

func getAllPosts(c *fiber.Ctx) error {
	return c.SendString("All Posts")
}

func getPostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Get Post by ID: " + postID)
}

func updatePostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Update Post by ID: " + postID)
}

func deletePostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Delete Post by ID: " + postID)
}

func createPost(c *fiber.Ctx) error {
	return c.SendString("Create Post")
}
