package service

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerService interface {
	Create(worker *models.Worker) (*models.Auth, error)
	Update(worker *models.Worker) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() models.Workers
}

type IAuthService interface {
	Create(firstName, lastName, role string) (*models.Auth, error)
	Update(auth *models.Auth) error
}

type IProjectService interface {
	Create(project *models.Project) (uuid.UUID, error)
	Get(clientID uuid.UUID) (models.Projects, error)
	Update(project *models.Project) error
	Delete(id uuid.UUID) error
}

type IThreadService interface {
	Create(thread *models.Thread) (uuid.UUID, error)
	Get(projectID uuid.UUID) (models.Threads, error)
	Update(thread *models.Thread) error
	Delete(id uuid.UUID) error
}

type ICampaignService interface {
	Update(camp *models.Campaign) error
}

type IClientService interface {
	Create(client *models.User) (uuid.UUID, error)
	Update(client *models.User) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.User, error)
	GetAll() (models.Users, error)
}
