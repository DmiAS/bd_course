package converters

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func convertUser(u models.User) ds.User {
	return ds.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		VkLink:    u.VkLink,
		TgLink:    u.TgLink,
	}
}
