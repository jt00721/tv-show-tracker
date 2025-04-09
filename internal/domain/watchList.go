package domain

import (
	"time"

	"gorm.io/gorm"
)

type Watchlist struct {
	ID                 uint           `gorm:"primaryKey"`
	ShowID             uint           `gorm:"not null" json:"show_id"`
	UserID             uint           `gorm:"not null" json:"user_id"`
	LastWatchedEpisode string         `json:"last_watched_episode"`
	EpisodesWatched    int            `json:"episodes_watched"`
	Status             string         `json:"status"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
