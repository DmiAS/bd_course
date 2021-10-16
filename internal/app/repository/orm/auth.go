package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a AuthRepository) GetAuth(login string) (*models.Auth, error) {
	auth := &models.Auth{}
	if err := a.db.Model(&models.Auth{}).Where("login = ?", login).First(auth).Error; err != nil {
		return nil, err
	}
	return auth, nil
}

func (a AuthRepository) Create(auth *models.Auth) error {
	return a.db.Create(auth).Error
}

func (a AuthRepository) Delete(id uuid.UUID) error {
	return a.db.Where("id = ?", id).Delete(&models.Auth{}).Error
}

func (a AuthRepository) Update(info *models.Auth) error {
	return a.db.Where("user_id = ?", info.UserID).Updates(info).Error
}
