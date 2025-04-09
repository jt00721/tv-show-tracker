package domain

import (
	"time"

	"gorm.io/gorm"
)

type Show struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"not null"`
	Summary       string `gorm:"not null"`
	ImageURL      string
	Rating        float64
	TotalEpisodes int
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
