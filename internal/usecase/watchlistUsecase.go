package usecase

import (
	"fmt"
	"log"

	"github.com/jt00721/tv-show-tracker/internal/domain"
	"github.com/jt00721/tv-show-tracker/internal/repository"
)

type WatchlistUsecase interface {
	CreateWatchlist(w *domain.Watchlist) error
	GetWatchlist(userID uint) ([]domain.Watchlist, error)
	GetWatchlistByID(id uint) (domain.Watchlist, error)
	UpdateWatchlist(w *domain.Watchlist) error
	DeleteWatchlist(watchlistID uint) error
	MarkWatched(watchlistID uint, lastWatchedEpisode string) error
}

type watchlistUsecase struct {
	Repo repository.WatchlistRepository
}

func NewWatchlistUsecase(r repository.WatchlistRepository) *watchlistUsecase {
	return &watchlistUsecase{Repo: r}
}

func (uc *watchlistUsecase) CreateWatchlist(w *domain.Watchlist) error {
	if w.ShowID <= 0 || w.UserID <= 0 {
		return fmt.Errorf("ShowID or UserID cannot be 0")
	}
	return uc.Repo.Create(w)
}

func (uc *watchlistUsecase) GetWatchlist(userID uint) ([]domain.Watchlist, error) {
	return uc.Repo.GetByUserID(userID)
}

func (uc *watchlistUsecase) GetWatchlistByID(id uint) (domain.Watchlist, error) {
	return uc.Repo.GetByID(id)
}

func (uc *watchlistUsecase) UpdateWatchlist(w *domain.Watchlist) error {
	return uc.Repo.Update(w)
}

func (uc *watchlistUsecase) DeleteWatchlist(watchlistID uint) error {
	return uc.Repo.Delete(watchlistID)
}

func (uc *watchlistUsecase) MarkWatched(watchlistID uint, lastWatchedEpisode string) error {
	watchlist, err := uc.GetWatchlistByID(watchlistID)
	if err != nil {
		log.Println("Error fetching watchlist for marking watched:", err)
		return fmt.Errorf("watchlist not found")
	}

	if lastWatchedEpisode == "" {
		return fmt.Errorf("last watched episode cannot be empty")
	}

	watchlist.LastWatchedEpisode = lastWatchedEpisode
	watchlist.EpisodesWatched++

	if watchlist.Status != "Watching" && watchlist.Status != "Completed" {
		watchlist.Status = "Watching"
	}

	if err := uc.Repo.Update(&watchlist); err != nil {
		log.Println("Error marking watchlist as watched:", err)
		return fmt.Errorf("failed to mark watchlist as watched")
	}

	log.Printf("Watchlist with ID(%d) marked as watched. Current Episodes Watched: %d", watchlistID, watchlist.EpisodesWatched)
	return nil
}
