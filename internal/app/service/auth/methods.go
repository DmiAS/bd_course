package auth

import (
	"encoding/base64"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const passwordSize = 12
const saltSize = 16

func (s *Service) Login(login, password string) (string, error) {
	rep := s.unit.GetAuthRepository()
	auth, err := rep.GetAuth(login)
	if err != nil {
		return "", err
	}
	if err := comparePassword(password, auth); err != nil {
		return "", err
	}

	return s.createToken(auth.UserID)
}

func (s *Service) GetRoleInfo(tokenStr string) (*models.IDs, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header[algHeader])
		}
		return s.signKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.Errorf("token %s is invalid", tokenStr)
	}

	return s.extractIdsFromToken(token)
}

func (s *Service) Create(firstName, lastName, role string) (*models.Auth, error) {
	var info *models.Auth
	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		auth := u.GetAuthRepository()
		id, err := auth.CreateIdRow(role)
		if err != nil {
			return err
		}

		password := gen.GenReadableString(passwordSize)
		login := gen.Login(firstName, lastName)

		encInfo, err := createAuthInfo(id, login, password)
		if err != nil {
			return err
		}
		if err := auth.Create(encInfo); err != nil {
			return err
		}
		info = &models.Auth{
			Login:    login,
			Password: password,
			UserID:   id,
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *Service) Update(info *models.Auth) error {
	encInfo, err := createAuthInfo(info.UserID, info.Login, info.Password)
	if err != nil {
		return err
	}

	auth := s.unit.GetAuthRepository()
	return auth.Update(encInfo)
}

func (s *Service) Delete(id uuid.UUID) error {
	ar := s.unit.GetAuthRepository()
	return ar.Delete(id)
}

func createAuthInfo(id uuid.UUID, login string, password string) (*models.Auth, error) {
	salt, err := gen.GenerateRandomString(saltSize)
	if err != nil {
		return nil, err
	}
	encP, err := gen.PasswordWithSalt([]byte(password), salt)
	if err != nil {
		return nil, err
	}
	auth := &models.Auth{
		Login:    login,
		Password: bytesToString(encP),
		Salt:     bytesToString(salt),
		UserID:   id,
	}

	return auth, nil
}

func bytesToString(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
