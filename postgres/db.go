package postgres

import (
	"fmt"
	"go_blog/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBConnection() {
	godotenv.Load()
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Post{})

	post := models.Post{Text: "hello world", Title: "post 1"}
	result := db.Create(&post)
	if result.Error != nil {
		log.Fatalf("Failed to create post: %v", result.Error)
	}

	// Read
	var firstPost models.Post

	if err := db.First(&firstPost).Error; err != nil {
		log.Fatalf("Failed to read the first post: %v", err)
	} else {
		fmt.Printf("Title of the first post: %s\n", firstPost.Title)
	}
	// Read
	// var product Post
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Post{Text: "hello world", Title: "post 1"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
