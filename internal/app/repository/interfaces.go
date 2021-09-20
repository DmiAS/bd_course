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

type IProjectRepository interface {
	Create(project *models.Project) (uuid.UUID, error)
	Get(clientID uuid.UUID) (models.Projects, error)
	Delete(id uuid.UUID) error
}
