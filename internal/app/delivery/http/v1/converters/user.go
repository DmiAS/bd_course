package converters

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertUserOutput(client models.User) ds.User {
	return ds.User{
		FirstName: client.FirstName,
		LastName:  client.LastName,
		VkLink:    client.VkLink,
		TgLink:    client.TgLink,
	}
}

func ConvertUserInput(u *ds.User) *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		VkLink:    u.VkLink,
		TgLink:    u.TgLink,
	}
}

func ConvertUserWitIDInput(u *ds.User, id uuid.UUID) *models.User {
	user := ConvertUserInput(u)
	user.ID = id
	return user
}
