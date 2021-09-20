package ds

import "github.com/google/uuid"

type Worker struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Grade     string `json:"grade"`
	Position  string `json:"position"`
}

type WorkerUUID struct {
	Worker
	UUID uuid.UUID `json:"uuid"`
}

type CreateWorkerInput struct {
	Worker
	Login string `json:"login"`
}

type CreateWorkerOutput struct {
	Login    string `json:"login"`
	Password string `json:"gen"`
}

type UpdateWorkerInput struct {
	Worker
}

type GetWorkerOutput struct {
	Worker
}

type GetAllWorkersOutput struct {
	Count   int          `json:"count"`
	Workers []WorkerUUID `json:"workers"`
}
