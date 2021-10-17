package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user *models.User) error {
	return u.db.Create(user).Error
}

func (u UserRepository) Update(user *models.User) error {
	return u.db.Updates(user).Error
}

func (u UserRepository) Get(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	res := u.db.Where("id = ?", id).First(user)
	return user, res.Error
}

func (u UserRepository) GetAll(role models.Role, created int64, limit int) models.Users {
	var users models.Users
	u.db.Limit(limit).
		Order("created desc").
		Where("role = ? and created < ?", role, created).
		Find(&users)
	return users
}
