package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IAuthService interface {
	Login(login, password string) (*models.RoleToken, error)
	GetRoleInfo(tokenStr string) (*models.UserInfo, error)
	Create(firstName, lastName string) (*models.Auth, error)
	Delete(id uuid.UUID) error
	Update(auth *models.Auth) error
}

type IWorkerService interface {
	Create(worker *models.WorkerEntity) (*models.Auth, error)
	Update(worker *models.WorkerEntity) error
	Get(id uuid.UUID) (*models.WorkerEntity, error)
	GetAll() *models.WorkersList
}

type IUserService interface {
	Create(user *models.User) (*models.Auth, error)
	Update(user *models.User) error
	Get(id uuid.UUID) (*models.User, error)
	GetAll(role models.Role) *models.UserList
}

type ICampaignService interface {
	GetAll(pagination *models.Pagination) (*models.CampaignsList, error)
	//methods to work with campaigns
	Get(id uuid.UUID) (*models.Campaign, error)
	GetCampaigns(targetologID uuid.UUID, pagination *models.Pagination) *models.CampaignsList
	// прикрепляет или открепляет кампанию к потоку
	AttachCampaign(threadID, campID uuid.UUID) error
	// назначает кампанию на таргетолога
	AssignCampaign(targetologID, campID uuid.UUID) error
}

type IProjectService interface {
	Create(clientID uuid.UUID, name string) error
	Get(projectID uuid.UUID) (*models.Project, error)
	GetAll(clientID uuid.UUID) *models.ProjectsList
	Update(project *models.Project) error
	Delete(projectID uuid.UUID) error
}

type IThreadService interface {
	Create(projectID uuid.UUID, name string) error
	Get(threadID uuid.UUID) (*models.Thread, error)
	GetAll(projectID uuid.UUID) *models.ThreadsList
	Update(thread *models.Thread) error
	Delete(threadID uuid.UUID) error
}

type IStatService interface {
	GetProjectStat(projectID uuid.UUID, from, to time.Time) (*models.ProjectStat, error)
	GetThreadStat(threadID uuid.UUID, from, to time.Time) (*models.ThreadStat, error)
	GetTargetologStat(targetologID uuid.UUID, from, to time.Time) (*models.TargetologStat, error)
}
