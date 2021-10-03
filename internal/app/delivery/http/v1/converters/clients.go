package converters

import (
	ds "github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func ConvertClientInput(client *ds.User) *models.Client {
	return &models.Client{User: *ConvertUserInput(client)}
}

func ConvertClientInputWithID(client *ds.User, id uuid.UUID) *models.Client {
	user := ConvertUserInput(client)
	user.ID = id
	return &models.Client{User: *user}
}

func convertClientOutput(client models.Client) ds.User {
	return ds.User{
		FirstName: client.User.FirstName,
		LastName:  client.User.LastName,
		VkLink:    client.User.VkLink,
		TgLink:    client.User.TgLink,
	}
}
func ConvertClientOutput(client *models.Client) *ds.User {
	user := convertClientOutput(*client)
	return &user
}

func ConvertGetAllClientsOutput(clients models.Clients) ds.GetAllClientsOutput {
	cs := make([]ds.ClientUUID, 0, len(clients))
	for _, client := range clients {
		user := convertClientOutput(client)
		cs = append(cs, ds.ClientUUID{
			User: user,
			UUID: client.User.ID,
		})
	}

	return ds.GetAllClientsOutput{
		Count:   len(cs),
		Clients: cs,
	}
}
