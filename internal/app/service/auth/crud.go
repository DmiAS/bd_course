package auth

import (
	"encoding/base64"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/google/uuid"
)

const passwordSize = 12
const saltSize = 16

type authInfo struct {
	login    string
	password []byte
	salt     []byte
	id       uuid.UUID
}

func (s *Service) Create(firstName, lastName, role string) (*models.Auth, error) {
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

	encInfo, err := encryptAuthInfo(authInfo)
	if err != nil {
		unit.Rollback()
		return nil, err
	}
	if err := auth.Create(encInfo); err != nil {
		unit.Rollback()
		return nil, err
	}
	unit.Commit()
	return &models.Auth{
		Login:    authInfo.login,
		Password: bytesToString(authInfo.password),
		UserID:   id,
	}, nil
}

func createAuthInfo(id uuid.UUID, firstName, lastName string) (*authInfo, error) {
	salt, err := gen.GenerateRandomString(saltSize)
	if err != nil {
		return nil, err
	}

	password, err := gen.GenerateRandomString(passwordSize)
	if err != nil {
		return nil, err
	}

	login := gen.Login(firstName, lastName)
	auth := &authInfo{
		login:    login,
		password: password,
		salt:     salt,
		id:       id,
	}

	return auth, nil
}

func encryptAuthInfo(info *authInfo) (*models.Auth, error) {
	hashedPassword, err := gen.PasswordWithSalt(info.password, info.salt)
	if err != nil {
		return nil, err
	}

	strSalt := bytesToString(info.salt)

	return &models.Auth{
		Login:    info.login,
		Password: hashedPassword,
		Salt:     strSalt,
		UserID:   info.id,
	}, nil
}

func bytesToString(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
