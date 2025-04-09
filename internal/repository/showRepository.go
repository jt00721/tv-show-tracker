package repository

import (
	"github.com/jt00721/tv-show-tracker/internal/domain"
	"gorm.io/gorm"
)

type ShowRepository struct {
	DB *gorm.DB
}

func (repo *ShowRepository) Create(show *domain.Show) error {
	return repo.DB.Create(show).Error
}

func (repo *ShowRepository) GetByID(id uint) (*domain.Show, error) {
	var show domain.Show
	err := repo.DB.First(&show, id).Error
	return &show, err
}

func (repo *ShowRepository) GetAll() ([]domain.Show, error) {
	var shows []domain.Show
	err := repo.DB.Find(&shows).Error
	return shows, err
}

func (repo *ShowRepository) Update(show *domain.Show) error {
	return repo.DB.Save(show).Error
}

func (repo *ShowRepository) Delete(id uint) error {
	return repo.DB.Delete(&domain.Show{}, id).Error
}
