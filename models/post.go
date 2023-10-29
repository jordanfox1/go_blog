package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Text  string `gorm:"unique"`
	Title string `gorm:"unique"`
}
