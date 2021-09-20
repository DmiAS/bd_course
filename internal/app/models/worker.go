package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Worker struct {
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Grade     string
	Position  string
	Cabs      pq.Int64Array
}
type Workers []Worker
