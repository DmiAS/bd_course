package models

import "github.com/google/uuid"

type Thread struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Name      string
}
type Threads []Thread
