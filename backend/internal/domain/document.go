package domain

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID                  uint           `gorm:"primarykey" json:"id"`
	UserID              uint           `gorm:"not null;index" json:"user_id"`
	User                User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Title               string         `gorm:"not null" json:"title"`
	Description         string         `json:"description"`
	FirebaseStoragePath string         `json:"firebase_storage_path"`
	FileType            string         `json:"file_type"`
	FileSize            int64          `json:"file_size"`
	Status              string         `gorm:"default:'active'" json:"status"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CreateDocumentDTO struct {
	UserID              uint   `json:"user_id" validate:"required"`
	Title               string `json:"title" validate:"required"`
	Description         string `json:"description"`
	FirebaseStoragePath string `json:"firebase_storage_path" validate:"required"`
	FileType            string `json:"file_type"`
	FileSize            int64  `json:"file_size"`
}

type UpdateDocumentDTO struct {
	Title       string `json:"title" validate:"omitempty"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"omitempty,oneof=active inactive archived"`
}
