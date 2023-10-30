package middlewares

import (
	"go_blog/models"
	"go_blog/validation"

	"github.com/gofiber/fiber/v2"
)

func ValidatePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post := &models.Post{} // Assuming you can access the data from the request body

		// Parse and validate the data
		if err := c.BodyParser(post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		// Validate the data using your validation package
		if err := validation.ValidatePost(post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Validation failed",
			})
		}

		// Data is valid, continue to the next handler
		return c.Next()
	}
}
