package auth

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go/v4"
)

const algHeader = "alg"

func comparePassword(password string, auth *models.Auth) error {
	bp := []byte(password)
	bs := []byte(auth.Salt)
	encP, err := gen.PasswordWithSalt(bp, bs)
	if err != nil {
		return err
	}
	if encP != auth.Password {
		return errors.New("password does not match")
	}
	return nil
}

type claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func (s *Service) createToken(id uuid.UUID) (string, error) {
	rep := s.unit.GetAuthRepository()
	role, err := rep.GetRole(id)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(s.expireDuration)),
			ID:        id.String(),
			IssuedAt:  jwt.At(time.Now()),
		},
	})
	return token.SignedString(s.signKey)
}

func (s *Service) extractIdsFromToken(token *jwt.Token) (*models.IDs, error) {
	claims, ok := token.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	return &models.IDs{
		ID:   id,
		Role: claims.Role,
	}, nil
}
