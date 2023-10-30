package storage

import (
	"fmt"
	"go_blog/models"
	"log"

	"gorm.io/gorm"
)

type PostgresStorage struct {
	DB *gorm.DB
}

func (p *PostgresStorage) GetPostByID(id int) (*models.Post, error) {
	post := &models.Post{}
	if err := p.DB.Where("id = ?", id).First(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostgresStorage) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	result := p.DB.Find(&posts)
	if result.Error != nil {
		log.Printf("unable to get all posts from DB - %v", result.Error)
	}
	return posts, nil
}

func (p *PostgresStorage) CreatePost(post *models.Post) error {
	if err := p.DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresStorage) UpdatePost(postID int, title, text string) error {
	existingPost, err := p.GetPostByID(postID)
	if err != nil {
		return err
	}

	existingPost.Text, existingPost.Title = text, title

	return p.DB.Save(existingPost).Error
}

func (p *PostgresStorage) DeletePost(id int) error {
	result := p.DB.Where("id = ?", id).Delete(&models.Post{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("post not found with ID: %d", id)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
