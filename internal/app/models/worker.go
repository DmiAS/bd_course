package models

import (
	"github.com/lib/pq"
)

type Worker struct {
	User
	Grade    string
	Position string
	Cabs     pq.Int64Array
}
type Workers []Worker
