package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerRepository interface {
	Create(worker *models.Worker) error
	Update(worker *models.Worker) error
	Get(id uuid.UUID) (*models.Worker, error)
	GetAll() models.Workers
}

type ICampaignRepository interface {
	GetAll() models.Campaigns
	GetCampaign(campaignID uuid.UUID) (*models.Campaign, error)
	GetCampaigns(workerID uuid.UUID) models.Campaigns
	GetThreadCampaigns(threadID uuid.UUID) models.Campaigns
	GetCampaignStat(campID uuid.UUID, from, to time.Time) []models.CampaignStat
	Update(camp *models.Campaign) error
}

type IAuthRepository interface {
	Create(auth *models.Auth) error
	CreateIdRow(role string) (uuid.UUID, error)
	Update(auth *models.Auth) error
	Delete(id uuid.UUID) error
}

type IProjectRepository interface {
	Create(project *models.Project) error
	Get(projectID uuid.UUID) (*models.Project, error)
	GetAll(clientID uuid.UUID) models.Projects
	Update(project *models.Project) error
	Delete(clientID, projectID uuid.UUID) error
}

type IThreadRepository interface {
	Create(thread *models.Thread) error
	Get(threadID uuid.UUID) (*models.Thread, error)
	GetAll(projectID uuid.UUID) models.Threads
	Update(thread *models.Thread) error
	Delete(projectID, threadID uuid.UUID) error
}

type IClientRepository interface {
	Create(client *models.Client) error
	Update(client *models.Client) error
	Get(id uuid.UUID) (*models.Client, error)
	GetAll() models.Clients
}
