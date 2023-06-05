package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
