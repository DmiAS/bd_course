package uwork

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/repository"
)

type UnitOfWork interface {
	WithRole(role models.Role) UnitOfWork
	WithTransaction(func(u UnitOfWork) error) error

	GetWorkerRepository() repository.IWorkerRepository
	GetAuthRepository() repository.IAuthRepository
	GetUserRepository() repository.IUserRepository
	GetProjectRepository() repository.IProjectRepository
	GetThreadsRepository() repository.IThreadRepository
	GetCampaignsRepository() repository.ICampaignRepository
}
