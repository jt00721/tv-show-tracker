package usecase

import (
	"github.com/jt00721/tv-show-tracker/internal/domain"
)

// MockWatchlistRepo satisfies the HabitRepository interface
type MockWatchlistRepo struct {
	CreateFn      func(*domain.Watchlist) error
	GetByUserIDFn func() ([]domain.Watchlist, error)
	GetByIDFn     func(uint) (domain.Watchlist, error)
	UpdateFn      func(*domain.Watchlist) error
	DeleteFn      func(uint) error
}

// Implement each method to call the corresponding function if set

func (m *MockWatchlistRepo) Create(h *domain.Watchlist) error {
	if m.CreateFn != nil {
		return m.CreateFn(h)
	}
	return nil
}

func (m *MockWatchlistRepo) GetByUserID(userID uint) ([]domain.Watchlist, error) {
	if m.GetByUserIDFn != nil {
		return m.GetByUserIDFn()
	}
	return nil, nil
}

func (m *MockWatchlistRepo) GetByID(id uint) (domain.Watchlist, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(id)
	}
	return domain.Watchlist{}, nil
}

func (m *MockWatchlistRepo) Update(h *domain.Watchlist) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(h)
	}
	return nil
}

func (m *MockWatchlistRepo) Delete(id uint) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}
