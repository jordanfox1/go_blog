package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Text  string `gorm:"unique;not null" validate:"required,min=1,max=25000"`
	Title string `gorm:"unique;not null" validate:"required,min=1,max=50"`
}
