package storage

import "go_blog/models"

type DataStorage interface {
	GetPostByID(id int) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(id int, title, text string) error
	DeletePost(id int) error
}
