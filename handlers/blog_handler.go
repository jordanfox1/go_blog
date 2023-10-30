package handlers

import "github.com/gofiber/fiber/v2"

type BlogHandlerImpl struct{}

func (h *BlogHandlerImpl) HandleGetAllPosts(c *fiber.Ctx) error {
	return c.SendString("All Posts")
}

func (h *BlogHandlerImpl) HandleGetPostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Get Post by ID: " + postID)
}

func (h *BlogHandlerImpl) HandleUpdatePostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Update Post by ID: " + postID)
}

func (h *BlogHandlerImpl) HandleDeletePostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Delete Post by ID: " + postID)
}

func (h *BlogHandlerImpl) HandleCreatePost(c *fiber.Ctx) error {
	return c.SendString("Create Post")
}
