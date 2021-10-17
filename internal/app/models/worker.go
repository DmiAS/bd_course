package models

import "github.com/google/uuid"

type WorkerEntity struct {
	User
	Grade    string `json:"grade"`
	Position string `json:"position"`
}

type Worker struct {
	UserID   uuid.UUID
	Grade    string
	Position string
}
type Workers []Worker

type WorkersList struct {
	Cursor  int64          `json:"cursor"`
	Amount  int            `json:"amount"`
	Workers []WorkerEntity `json:"workers"`
}
