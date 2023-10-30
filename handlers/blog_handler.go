package handlers

import (
	"fmt"
	"go_blog/models"
	"go_blog/storage"
	"log"
	"strconv"
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
	postID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Printf("invalid id, could not convert to int error: %v", err)
		return c.SendStatus(400)
	}

	post, err := h.Storage.GetPostByID(postID)
	if err != nil {
		fmt.Printf("could not get post with id %v. error %v", postID, err)
		return c.SendStatus(400)
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("Title: %s, Text: %s\n", post.Title, post.Text))

	return c.SendString(response.String())
}

func (h *BlogHandlerImpl) HandleUpdatePostByID(c *fiber.Ctx) error {
	var newPost models.Post

	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	postID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error converting id to int",
		})
	}

	err = h.Storage.UpdatePost(postID, newPost.Title, newPost.Text)
	if err != nil {
		return err
	}
	return c.SendString(fmt.Sprintf("Updated Post with ID: %v", postID))
}

func (h *BlogHandlerImpl) HandleDeletePostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	return c.SendString("Delete Post by ID: " + postID)
}

func (h *BlogHandlerImpl) HandleCreatePost(c *fiber.Ctx) error {
	// Initialize an empty Post model
	var post models.Post

	// Parse and decode the request body into the Post model
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Now 'post' contains the data from the request body

	// Call the CreatePost method from the storage layer
	err := h.Storage.CreatePost(&post)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
