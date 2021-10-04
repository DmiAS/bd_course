package service

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerService interface {
	Create(worker *models.Worker) (*models.Auth, error)
	Update(worker *models.Worker) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() models.Workers
}

type IClientService interface {
	Create(client *models.Client) (*models.Auth, error)
	Update(client *models.Client) error
	Get(id uuid.UUID) (*models.Client, error)
	GetAll() models.Clients
}

type IAuthService interface {
	Create(firstName, lastName, role string) (*models.Auth, error)
	Delete(id uuid.UUID) error
	Update(auth *models.Auth) error
}

type IProjectService interface {
	Create(project *models.Project) (uuid.UUID, error)
	Get(clientID, projectID uuid.UUID) (*models.Project, error)
	GetAll() models.Projects
	Update(project *models.Project) error
	Delete(id uuid.UUID) error
}

type IThreadService interface {
	Create(thread *models.Thread) (uuid.UUID, error)
	Get(projectID uuid.UUID) (models.Threads, error)
	Update(thread *models.Thread) error
	Delete(id uuid.UUID) error
}
