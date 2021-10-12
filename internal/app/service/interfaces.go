package service

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IAuthService interface {
	Create(firstName, lastName, role string) (*models.Auth, error)
	Delete(id uuid.UUID) error
	Update(auth *models.Auth) error
}

type IWorkerService interface {
	Create(worker *models.Worker) (*models.Auth, error)
	Update(worker *models.Worker) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() *models.WorkersList
}

type ICampaignService interface {
	GetAll() *models.CampaignsList
	//methods to work with campaigns
	GetCampaigns(id uuid.UUID) *models.CampaignsList
	// прикрепляет или открепляет кампанию к потоку
	AttachCampaign(threadID, campID uuid.UUID) error
	// назначает кампанию на таргетолога
	AssignCampaign(targetologID, campID uuid.UUID) error
}

type IClientService interface {
	Create(client *models.Client) (*models.Auth, error)
	Update(client *models.Client) error
	Get(id uuid.UUID) (*models.Client, error)
	GetAll() *models.ClientsList
}

type IProjectService interface {
	Create(project *models.Project) error
	Get(clientID, projectID uuid.UUID) (*models.Project, error)
	GetAll(clientID uuid.UUID) *models.ProjectsList
	Update(project *models.Project) error
	Delete(clientID, projectID uuid.UUID) error
}

type IThreadService interface {
	Create(thread *models.Thread) error
	Get(projectID, threadID uuid.UUID) (*models.Thread, error)
	GetAll(projectID uuid.UUID) *models.ThreadsList
	Update(thread *models.Thread) error
	Delete(projectID, threadID uuid.UUID) error
}
