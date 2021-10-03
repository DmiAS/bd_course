package auth

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/google/uuid"
)

const passwordSize = 12
const saltSize = 16

func (s *Service) CreateAuth(firstName, lastName, role string) (*models.Auth, error) {
	unit := s.unit.WithTransaction()
	auth := unit.GetAuthRepository()

	id, err := auth.CreateIdRow(role)
	if err != nil {
		unit.Rollback()
		return nil, err
	}

	authInfo, err := createAuthInfo(id, firstName, lastName)
	if err != nil {
		unit.Rollback()
		return nil, err
	}

	if err := auth.Create(authInfo); err != nil {
		unit.Rollback()
		return nil, err
	}
	unit.Commit()
	return authInfo, nil
}

func createAuthInfo(id uuid.UUID, firstName, lastName string) (*models.Auth, error) {
	salt, err := gen.GenerateRandomString(saltSize)
	if err != nil {
		return nil, err
	}

	password, err := gen.GenerateRandomString(passwordSize)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := gen.PasswordWithSalt(password, salt)
	if err != nil {
		return nil, err
	}

	login := gen.Login(firstName, lastName)
	auth := &models.Auth{
		Login:    login,
		Password: hashedPassword,
		Salt:     string(salt),
		ID:       id,
	}

	return auth, nil
}
