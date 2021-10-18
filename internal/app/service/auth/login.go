package auth

import (
	"encoding/base64"
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/pkg/gen"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/golang-jwt/jwt/v4"
)

const algHeader = "alg"

func comparePassword(password string, auth models.Auth) error {
	bp := []byte(password)
	bs, err := base64.URLEncoding.DecodeString(auth.Salt)
	if err != nil {
		return err
	}
	encP, err := gen.PasswordWithSalt(bp, bs)
	if err != nil {
		return err
	}

	pass := base64.URLEncoding.EncodeToString(encP)
	if pass != auth.Password {
		return errors.New("password does not match")
	}
	return nil
}

type claims struct {
	Role models.Role `json:"role"`
	jwt.RegisteredClaims
}

func (s *Service) createToken(id uuid.UUID, role models.Role) (*models.RoleToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expireDuration)),
			ID:        id.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	jtoken, err := token.SignedString(s.signKey)
	if err != nil {
		return nil, err
	}
	return &models.RoleToken{
		Token: jtoken,
		Role:  role,
	}, nil
}

func (s *Service) extractIdsFromToken(token *jwt.Token) (*models.UserInfo, error) {
	claims, ok := token.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	return &models.UserInfo{
		ID:   id,
		Role: claims.Role,
	}, nil
}
