package converters

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertGetClientOutput(client *models.User) *ds.User {
	user := ConvertUserOutput(*client)
	return &user
}
func ConvertGetAllClientOutput(clients models.Users) *ds.GetAllClientsOutput {
	cnt := len(clients)
	cs := make([]ds.ClientUUID, 0, cnt)

	for _, c := range clients {
		cs = append(cs, ds.ClientUUID{
			User: ConvertUserOutput(c),
			UUID: c.ID,
		})
	}

	return &ds.GetAllClientsOutput{
		Count:   cnt,
		Clients: cs,
	}
}
