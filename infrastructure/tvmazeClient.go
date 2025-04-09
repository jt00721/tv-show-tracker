package infrastructure

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jt00721/tv-show-tracker/internal/domain"
)

type TVMazeClient struct {
	client *resty.Client
	apiURL string
}

func NewTVMazeClient(apiURL string) *TVMazeClient {
	return &TVMazeClient{
		client: resty.New(),
		apiURL: apiURL,
	}
}

// tvmazeClient.go
type apiShow struct {
	Show struct {
		Name    string `json:"name"`
		Summary string `json:"summary"`
		Rating  struct {
			Average float64 `json:"average"`
		} `json:"rating"`
		Image struct {
			Medium string `json:"medium"`
		} `json:"image"`
	} `json:"show"`
}

func (c *TVMazeClient) SearchShows(query string) ([]domain.Show, error) {
	url := fmt.Sprintf("%s/search/shows?q=%s", c.apiURL, query)

	var apiResults []apiShow
	resp, err := c.client.R().
		SetResult(&apiResults).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("TV Maze API error: %s", resp.Status())
	}

	var shows []domain.Show
	for _, r := range apiResults {
		shows = append(shows, domain.Show{
			Title:    r.Show.Name,
			Summary:  r.Show.Summary,
			Rating:   r.Show.Rating.Average,
			ImageURL: r.Show.Image.Medium,
		})
	}

	return shows, nil
}
