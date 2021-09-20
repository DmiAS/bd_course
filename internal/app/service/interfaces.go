package service

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerService interface {
	Create(worker *models.Worker) (uuid.UUID, error)
	Update(worker *models.Worker) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() (models.Workers, error)
}

type IAuthService interface {
	Create(id uuid.UUID, login string) (string, error)
}
