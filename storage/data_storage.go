package storage

import "go_blog/models"

type DataStorage interface {
	GetPostByID(id int) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(post models.Post) error
	DeletePost(id int) error
}
