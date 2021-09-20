package models

import "github.com/google/uuid"

type Project struct {
	ID       uuid.UUID
	ClientID uuid.UUID
	Name     string
}
type Projects []Project
