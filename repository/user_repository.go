package repository

import (
	"gorm.io/gorm"

	"github.com/khasyah-fr/phinco-attendance/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	result := ur.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	result := ur.db.Where("username = ?", username).First(user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) UpdatePasswordByUsername(username string, newPassword string) error {
	user := &models.User{}
	result := ur.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	result = ur.db.Model(&user).Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
