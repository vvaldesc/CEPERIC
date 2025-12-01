package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Email       string         `gorm:"uniqueIndex;not null" json:"email"`
	Name        string         `gorm:"not null" json:"name"`
	FirebaseUID string         `gorm:"uniqueIndex" json:"firebase_uid,omitempty"`
	Role        string         `gorm:"default:'user'" json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CreateUserDTO struct {
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required,min=2"`
	FirebaseUID string `json:"firebase_uid"`
	Role        string `json:"role" validate:"omitempty,oneof=admin user"`
}

type UpdateUserDTO struct {
	Email string `json:"email" validate:"omitempty,email"`
	Name  string `json:"name" validate:"omitempty,min=2"`
	Role  string `json:"role" validate:"omitempty,oneof=admin user"`
}
