package ds

import "github.com/google/uuid"

type Worker struct {
	User
	Grade    string `json:"grade"`
	Position string `json:"position"`
}

type WorkerUUID struct {
	Worker
	UUID uuid.UUID `json:"uuid"`
}

type GetAllWorkersOutput struct {
	Count   int          `json:"count"`
	Workers []WorkerUUID `json:"workers"`
}
