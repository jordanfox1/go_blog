package handlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	HandleGetAllPosts(c *fiber.Ctx) error
	HandleGetPostByID(c *fiber.Ctx) error
	HandleUpdatePostByID(c *fiber.Ctx) error
	HandleDeletePostByID(c *fiber.Ctx) error
	HandleCreatePost(c *fiber.Ctx) error
}
