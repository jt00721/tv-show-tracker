package repository

import (
	"github.com/jt00721/tv-show-tracker/internal/domain"
	"gorm.io/gorm"
)

type WatchlistRepository interface {
	Create(w *domain.Watchlist) error
	GetByUserID(userID uint) ([]domain.Watchlist, error)
	GetByID(id uint) (domain.Watchlist, error)
	Update(w *domain.Watchlist) error
	Delete(watchlistID uint) error
}

type watchlistRepository struct {
	DB *gorm.DB
}

func NewWatchlistRepository(DB *gorm.DB) *watchlistRepository {
	return &watchlistRepository{DB: DB}
}

func (r *watchlistRepository) Create(w *domain.Watchlist) error {
	return r.DB.Create(w).Error
}

func (r *watchlistRepository) GetByUserID(userID uint) ([]domain.Watchlist, error) {
	var watchlists []domain.Watchlist
	err := r.DB.Where("user_id = ?", userID).Find(&watchlists).Error
	return watchlists, err
}

func (r *watchlistRepository) GetByID(id uint) (domain.Watchlist, error) {
	var watchlist domain.Watchlist
	err := r.DB.First(&watchlist, id).Error
	return watchlist, err
}

func (r *watchlistRepository) Update(w *domain.Watchlist) error {
	return r.DB.Save(w).Error
}

func (r *watchlistRepository) Delete(watchlistID uint) error {
	return r.DB.Delete(&domain.Watchlist{}, watchlistID).Error
}
