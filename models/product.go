package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int            `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"default:current_timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"not null"`
	Status      string         `gorm:"not null"`
	VariantsIds []string       `gorm:"type:json"` // We use JSON for the 'variants_ids' field, which is an array of strings
	SeoList     []string       `gorm:"type:json"` // Store the values separated by commas or any other delimiter
	BrandID     string         `gorm:"not null"`
	ImageID     string         `gorm:"not null"`
	ImagesIds   []string       `gorm:"type:json"` // We use JSON for the 'images_ids' field, which is an array of strings
	Description string         `gorm:"not null"`
}
