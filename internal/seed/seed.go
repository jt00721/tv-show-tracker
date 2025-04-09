package seed

import (
	"log"

	"github.com/jt00721/tv-show-tracker/internal/domain"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	shows := []domain.Show{
		{Title: "Stranger Things", Summary: "Mystery in Hawkins", Rating: 8.7, ImageURL: "https://example.com/stranger.jpg", TotalEpisodes: 25},
		{Title: "Breaking Bad", Summary: "A chemistry teacher turns meth cook", Rating: 9.5, ImageURL: "https://example.com/breaking.jpg", TotalEpisodes: 62},
		{Title: "The Office", Summary: "Everyday office life", Rating: 8.9, ImageURL: "https://example.com/office.jpg", TotalEpisodes: 201},
	}

	for _, show := range shows {
		if err := db.Create(&show).Error; err != nil {
			log.Println("Failed to seed show:", show.Title, "Error:", err)
			return err
		}
	}

	log.Println("Seeded initial shows successfully")
	return nil
}
