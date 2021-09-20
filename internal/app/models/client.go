package models

import "github.com/google/uuid"

type Client struct {
	UUID      uuid.UUID
	FirstName string
	LastName  string
}
type Clients []Client
