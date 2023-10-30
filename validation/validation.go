package validation

import (
	"go_blog/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidatePost(post *models.Post) error {
	return validate.Struct(post)
}
