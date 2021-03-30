package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `validate:"min=5,max=28"`
	Password string `validate:"min=5,max=150"`
	Email    string `validate:"required,email"`
	Posts    []Post
}

type Post struct {
	gorm.Model
	Title   string `validate:"min=5,max=28"`
	Content string `validate:"min=5,max=500"`
	UserID  uint
}
