package models

import "github.com/google/uuid"

const (
	AdminRole  = "admin"
	WorkerRole = "worker"
	ClientRole = "client"
)

type IDs struct {
	ID   uuid.UUID
	Role string
}

type Auth struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Salt     string    `json:"-"`
	UserID   uuid.UUID `json:"-"`
}
