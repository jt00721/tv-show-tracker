package usecase

import (
	"github.com/jt00721/tv-show-tracker/infrastructure"
	"github.com/jt00721/tv-show-tracker/internal/domain"
)

type ShowUsecase interface {
	SearchShows(query string) ([]domain.Show, error)
}

type showUsecase struct {
	tvmazeClient *infrastructure.TVMazeClient
}

func NewShowUsecase(client *infrastructure.TVMazeClient) *showUsecase {
	return &showUsecase{tvmazeClient: client}
}

func (uc *showUsecase) SearchShows(query string) ([]domain.Show, error) {
	return uc.tvmazeClient.SearchShows(query)
}
