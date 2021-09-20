package auth

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
)

const passwordSize = 12
const saltSize = 16

func (s *Service) CreateAuth(id uuid.UUID, firstName, lastName string) (*models.Auth, error) {

	salt, err := gen.GenerateRandomString(saltSize)
	if err != nil {
		return nil, err
	}

	password, err := gen.GenerateRandomString(passwordSize)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := gen.GenPasswordWithSalt(password, salt)
	if err != nil {
		return nil, err
	}

	login := gen.GenLogin(firstName, lastName)
	auth := &models.Auth{
		ID:       id,
		Login:    login,
		Password: hashedPassword,
	}

	if err := s.rep.Create(auth); err != nil {
		return nil, err
	}

	return auth, nil
}
