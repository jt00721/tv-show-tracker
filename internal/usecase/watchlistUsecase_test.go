package usecase_test

import (
	"errors"
	"testing"

	"github.com/jt00721/tv-show-tracker/internal/domain"
	"github.com/jt00721/tv-show-tracker/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestMarkCompleted(t *testing.T) {
	tests := []struct {
		name                  string
		watchlistID           uint
		lastWatchedEpisode    string
		mockGetByID           func(uint) (domain.Watchlist, error)
		mockUpdate            func(*domain.Watchlist) error
		wantErr               bool
		errContains           string
		expectEpisodesWatched int
	}{
		{
			name:               "first time completion",
			watchlistID:        1,
			lastWatchedEpisode: "S01E01",
			mockGetByID: func(id uint) (domain.Watchlist, error) {
				return domain.Watchlist{ID: id, ShowID: 1, UserID: 1, LastWatchedEpisode: "", EpisodesWatched: 0, Status: ""}, nil
			},
			mockUpdate: func(w *domain.Watchlist) error {
				assert.Equal(t, 1, w.EpisodesWatched)
				assert.Equal(t, "S01E01", w.LastWatchedEpisode)
				return nil
			},
			wantErr:               false,
			expectEpisodesWatched: 1,
		},
		{
			name:               "watchlist in watching status",
			watchlistID:        2,
			lastWatchedEpisode: "S01E02",
			mockGetByID: func(id uint) (domain.Watchlist, error) {
				return domain.Watchlist{ID: id, ShowID: 1, UserID: 1, LastWatchedEpisode: "S01E01", EpisodesWatched: 1, Status: "Watching"}, nil
			},
			mockUpdate: func(w *domain.Watchlist) error {
				assert.Equal(t, 2, w.EpisodesWatched)
				assert.Equal(t, "S01E02", w.LastWatchedEpisode)
				return nil
			},
			wantErr:               false,
			expectEpisodesWatched: 2,
		},
		{
			name:               "watchlist not found",
			watchlistID:        3,
			lastWatchedEpisode: "S01E03",
			mockGetByID: func(id uint) (domain.Watchlist, error) {
				return domain.Watchlist{}, errors.New("watchlist not found")
			},
			mockUpdate:  nil,
			wantErr:     true,
			errContains: "watchlist not found",
		},
		{
			name:               "update fails",
			watchlistID:        4,
			lastWatchedEpisode: "S01E04",
			mockGetByID: func(id uint) (domain.Watchlist, error) {
				return domain.Watchlist{ID: id, ShowID: 1, UserID: 1, LastWatchedEpisode: "S01E02", EpisodesWatched: 2, Status: "Watching"}, nil
			},
			mockUpdate: func(w *domain.Watchlist) error {
				return errors.New("db error")
			},
			wantErr:     true,
			errContains: "failed to mark watchlist as watched",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockWatchlistRepo{
				GetByIDFn: tt.mockGetByID,
				UpdateFn:  tt.mockUpdate,
			}
			uc := usecase.NewWatchlistUsecase(mockRepo)

			err := uc.MarkWatched(tt.watchlistID, tt.lastWatchedEpisode)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
