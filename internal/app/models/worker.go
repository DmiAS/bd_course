package models

import (
	"github.com/google/uuid"
)

type Worker struct {
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Grade     string
	Position  string
}
type Workers []Worker
