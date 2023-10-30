package handlers

import (
	"fmt"
	"go_blog/models"
	"go_blog/storage"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BlogHandlerImpl struct {
	Storage storage.DataStorage
}

func (h *BlogHandlerImpl) HandleGetAllPosts(c *fiber.Ctx) error {
	posts, err := h.Storage.GetAllPosts()
	if err != nil {
		log.Fatalf("Failed to retrieve posts: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve posts")
	}

	// Initialize a buffer to build the response
	var response strings.Builder

	// Iterate over the posts and build the response
	for _, post := range posts {
		response.WriteString(fmt.Sprintf("Title: %s, Text: %s\n", post.Title, post.Text))
	}

	return c.SendString(response.String())
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
	post := models.Post{Text: "hello world", Title: "post 1"}
	err := h.Storage.CreatePost(&post)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	return c.SendStatus(200)
}
