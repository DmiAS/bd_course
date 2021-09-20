package repository

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerRepository interface {
	Create(worker *models.Worker) (uuid.UUID, error)
	Update(worker *models.Worker) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() (models.Workers, error)
}

type IAuthRepository interface {
	Create(auth *models.Auth) error
}
