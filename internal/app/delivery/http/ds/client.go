package ds

import "github.com/google/uuid"

type Client struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ClientUUID struct {
	Client
	UUID uuid.UUID `json:"uuid"`
}

type CreateClientInput struct {
	Client
	Login string `json:"login"`
}

type CreateClientOutput struct {
	Login    string `json:"login"`
	Password string `json:"gen"`
}

type UpdateClientInput struct {
	Client
}

type GetClientOutput struct {
	Client
}

type GetAllClientsOutput struct {
	Count   int          `json:"count"`
	Clients []ClientUUID `json:"clients"`
}
