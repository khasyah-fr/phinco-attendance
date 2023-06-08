package repository

import (
	"gorm.io/gorm"

	"github.com/khasyah-fr/phinco-attendance/models"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func (lr *LocationRepository) GetAllLocations() ([]models.Location, error) {
	var locations []models.Location
	result := lr.db.Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}
	return locations, nil
}
