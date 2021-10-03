package converters

import (
	ds2 "github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertGetClientOutput(client *models.User) *ds2.User {
	user := ConvertUserOutput(*client)
	return &user
}
func ConvertGetAllClientOutput(clients models.Users) *ds2.GetAllClientsOutput {
	cnt := len(clients)
	cs := make([]ds2.ClientUUID, 0, cnt)

	for _, c := range clients {
		cs = append(cs, ds2.ClientUUID{
			User: ConvertUserOutput(c),
			UUID: c.ID,
		})
	}

	return &ds2.GetAllClientsOutput{
		Count:   cnt,
		Clients: cs,
	}
}
