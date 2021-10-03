package repository

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerRepository interface {
	Create(worker *models.Worker) error
	Update(worker *models.Worker) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() models.Workers
}

type IAuthRepository interface {
	Create(auth *models.Auth) error
	CreateIdRow(role string) (uuid.UUID, error)
	Delete(id uuid.UUID) error
}

type IProjectRepository interface {
	Create(project *models.Project) error
	Get(clientID uuid.UUID) (models.Projects, error)
	Update(project *models.Project) error
	Delete(id uuid.UUID) error
}

type IThreadRepository interface {
	Create(thread *models.Thread) (uuid.UUID, error)
	Get(projectID uuid.UUID) (models.Threads, error)
	Update(thread *models.Thread) error
	Delete(id uuid.UUID) error
}

type ICampaignRepository interface {
	UpdateWorker(campID, workerID uuid.UUID) error
	UpdateThread(threadID uuid.UUID) error
}

//type IClientRepository interface {
//	Create(client *models.Client) (uuid.UUID, error)
//	Update(client *models.Client) error
//	Delete(id uuid.UUID) error
//	Get(id uuid.UUID) (*models.Client, error)
//	GetAll() (models.Clients, error)
//}
