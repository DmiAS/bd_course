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
	Update(info *models.Auth, userID uuid.UUID, role models.Role) error
}

type IWorkerService interface {
	Create(worker *models.WorkerEntity) (*models.Auth, error)
	Update(worker *models.WorkerEntity) error
	Get(id uuid.UUID) (*models.WorkerEntity, error)
	GetAll(pagination *models.Pagination) *models.WorkersList
}

type IUserService interface {
	Create(user *models.User) (*models.Auth, error)
	Update(user *models.User) error
	Get(id uuid.UUID) (*models.User, error)
	GetAll(role models.Role, pag *models.Pagination) *models.UserList
}

type ICampaignService interface {
	GetAll(pagination *models.Pagination) (*models.CampaignsList, error)
	//methods to work with campaigns
	Get(id uuid.UUID) (*models.Campaign, error)
	GetThreadCampaigns(threadID uuid.UUID, pagination *models.Pagination) *models.CampaignsList
	GetCampaigns(targetologID uuid.UUID, pagination *models.Pagination) *models.CampaignsList
	// прикрепляет или открепляет кампанию к потоку
	AttachCampaign(threadID, campID uuid.UUID) error
	// назначает кампанию на таргетолога
	AssignCampaign(targetologID, campID uuid.UUID) error
}

type IProjectService interface {
	Create(clientID uuid.UUID, name string) error
	Get(projectID, userID uuid.UUID, role models.Role) (*models.Project, error)
	GetAll(clientID uuid.UUID, pag *models.Pagination) *models.ProjectsList
	Update(project *models.Project) error
	Delete(projectID uuid.UUID) error
}

type IThreadService interface {
	Create(projectID uuid.UUID, name string) error
	Get(threadID uuid.UUID, userID uuid.UUID, role models.Role) (*models.Thread, error)
	GetAll(projectID uuid.UUID, userID uuid.UUID, role models.Role, pagination *models.Pagination) (*models.ThreadsList, error)
	Update(thread *models.Thread) error
	Delete(threadID uuid.UUID) error
}

type IStatService interface {
	GetFullCampaignStat(campaignID, userID uuid.UUID, role models.Role, from, to time.Time) (*models.CampStat, error)
	GetProjectStat(projectID, userID uuid.UUID, role models.Role, from, to time.Time) (*models.ProjectStat, error)
	GetThreadStat(threadID, userID uuid.UUID, role models.Role, from, to time.Time) (*models.ThreadStat, error)
	GetTargetologStat(targetologID uuid.UUID, from, to time.Time) (*models.TargetologStat, error)
}
