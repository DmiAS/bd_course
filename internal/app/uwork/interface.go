package uwork

import "github.com/DmiAS/bd_course/internal/app/repository"

type Role = int

const (
	Targetolog Role = iota
	Admin
	Client
)

type UnitOfWork interface {
	WithRole(role Role) UnitOfWork
	WithTransaction(func(u UnitOfWork) error) error

	GetWorkerRepository() repository.IWorkerRepository
	GetAuthRepository() repository.IAuthRepository
	GetClientRepository() repository.IClientRepository
	GetProjectRepository() repository.IProjectRepository
	GetThreadsRepository() repository.IThreadRepository
	GetCampaignsRepository() repository.ICampaignRepository
}
