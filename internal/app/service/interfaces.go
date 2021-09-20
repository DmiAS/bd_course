package service

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerService interface {
	CreateWorker(worker *models.Worker) (uuid.UUID, error)
	UpdateWorker(worker *models.Worker) error
}

type IAuthService interface {
	CreateAuth(id uuid.UUID, login string) (string, error)
}
