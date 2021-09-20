package models

import "github.com/google/uuid"

type Auth struct {
	UUID     uuid.UUID
	Login    string
	Password string
}
