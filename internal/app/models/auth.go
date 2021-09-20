package models

import "github.com/google/uuid"

type Auth struct {
	ID       uuid.UUID
	Login    string
	Password string
}
