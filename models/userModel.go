package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// username, email, password
	Username *string `json:"username" gorm:"unique"`
	Email    *string `json:"email" gorm:"unique"`
	Password string  `json:"password"`
}

type UserRequest struct {
	Username *string `json:"username" binding:"required"`
	Email    *string `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	// Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
