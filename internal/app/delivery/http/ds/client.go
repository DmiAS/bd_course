package ds

import "github.com/google/uuid"

type ClientUUID struct {
	User
	UUID uuid.UUID `json:"uuid"`
}

type GetAllClientsOutput struct {
	Count   int          `json:"count"`
	Clients []ClientUUID `json:"clients"`
}
