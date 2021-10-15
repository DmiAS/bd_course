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

func (a AuthRepository) GetRole(id uuid.UUID) (string, error) {
	var role string
	if err := a.db.Model(&models.IDs{}).
		Select("role").
		Where("id = ?", id).
		Scan(&role).
		Error; err != nil {
		return "", err
	}
	return role, nil
}

func (a AuthRepository) CreateIdRow(role string) (uuid.UUID, error) {
	id := models.IDs{ID: uuid.New(), Role: role}
	result := a.db.Create(&id)

	if err := result.Error; err != nil {
		return uuid.UUID{}, err
	}

	return id.ID, nil
}

func (a AuthRepository) Create(auth *models.Auth) error {
	return a.db.Create(auth).Error
}

func (a AuthRepository) Delete(id uuid.UUID) error {
	return a.db.Delete(&models.IDs{ID: id}).Error
}

func (a AuthRepository) Update(info *models.Auth) error {
	return a.db.Where("user_id = ?", info.UserID).Updates(info).Error
}
