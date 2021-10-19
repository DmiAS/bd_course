package auth

import (
	"encoding/base64"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const passwordSize = 12
const saltSize = 16

func (s *Service) Login(login, password string) (*models.RoleToken, error) {
	rep := s.unit.GetAuthRepository()
	auth, err := rep.GetAuthWithRole(login)
	if err != nil {
		return nil, err
	}
	if err := comparePassword(password, auth.Auth); err != nil {
		return nil, err
	}

	return s.createToken(auth.Auth.UserID, auth.Role)
}

func (s *Service) GetRoleInfo(tokenStr string) (*models.UserInfo, error) {
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

func (s *Service) Create(firstName, lastName string) (*models.Auth, error) {
	info := &models.Auth{
		Login:    gen.Login(firstName, lastName),
		Password: gen.GenReadableString(passwordSize),
		UserID:   uuid.New(),
	}

	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		auth := u.GetAuthRepository()
		encInfo, err := encryptInfo(info)
		if err != nil {
			return err
		}
		if err := auth.Create(encInfo); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *Service) Update(info *models.Auth, userID uuid.UUID, role models.Role) error {
	rep := s.unit.GetUserRepository()
	user, err := rep.Get(info.UserID)
	if err != nil {
		return errors.New("access denied")
	}
	if (user.Role == models.AdminRole && role == models.AdminRole) || (role == models.ClientRole || role == models.WorkerRole) {
		if userID != user.ID {
			return errors.New("access denied")
		}
	}
	encInfo, err := encryptInfo(info)
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

func encryptInfo(info *models.Auth) (*models.Auth, error) {
	salt, err := gen.GenerateRandomString(saltSize)
	if err != nil {
		return nil, err
	}
	encP, err := gen.PasswordWithSalt([]byte(info.Password), salt)

	if err != nil {
		return nil, err
	}
	auth := &models.Auth{
		Login:    info.Login,
		Password: bytesToString(encP),
		Salt:     bytesToString(salt),
		UserID:   info.UserID,
	}
	return auth, nil
}

func bytesToString(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
