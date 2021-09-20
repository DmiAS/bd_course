package auth

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
)

const passwordSize = 12
const saltSize = 16

func (s *Service) CreateAuth(id uuid.UUID, login string) (string, error) {
	salt, err := gen.GenerateRandomString(saltSize)
	if err != nil {
		return "", err
	}

	password, err := gen.GenerateRandomString(passwordSize)
	if err != nil {
		return "", err
	}

	hashedPassword, err := gen.GenPasswordWithSalt(password, salt)
	if err != nil {
		return "", err
	}

	auth := &models.Auth{
		UUID:     id,
		Login:    login,
		Password: hashedPassword,
	}
	return string(password), s.rep.CreateAuth(auth)
}
