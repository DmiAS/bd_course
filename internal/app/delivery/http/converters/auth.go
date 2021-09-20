package converters

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertAuthOutput(auth *models.Auth) *ds.Auth {
	return &ds.Auth{
		Login:    auth.Login,
		Password: auth.Password,
	}
}
