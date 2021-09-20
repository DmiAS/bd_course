package converters

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertCreateClientInput(req *ds.CreateClientInput) (*models.Client, string) {
	return &models.Client{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}, req.Login
}

func ConvertCreateClientOutput(login, password string) *ds.CreateClientOutput {
	return &ds.CreateClientOutput{
		Login:    login,
		Password: password,
	}
}

func ConvertUpdateClientInput(client *ds.UpdateClientInput, id uuid.UUID) *models.Client {
	return &models.Client{
		UUID:      id,
		FirstName: client.FirstName,
		LastName:  client.LastName,
	}
}

func convertToClient(client models.Client) ds.Client {
	return ds.Client{
		FirstName: client.FirstName,
		LastName:  client.LastName,
	}
}

func ConvertGetClientOutput(client *models.Client) *ds.GetClientOutput {
	return &ds.GetClientOutput{convertToClient(*client)}
}

func ConvertGetAllClientOutput(clients models.Clients) *ds.GetAllClientsOutput {
	cnt := len(clients)
	cs := make([]ds.ClientUUID, 0, cnt)

	for _, c := range clients {
		cs = append(cs, ds.ClientUUID{
			Client: convertToClient(c),
			UUID:   c.UUID,
		})
	}

	return &ds.GetAllClientsOutput{
		Count:   cnt,
		Clients: cs,
	}
}
