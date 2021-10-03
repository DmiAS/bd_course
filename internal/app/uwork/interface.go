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
	WithTransaction() UnitOfWork
	Commit()
	Rollback()
	GetClientRepository() repository.IClientRepository
	GetWorkerRepository() repository.IWorkerRepository
	GetAuthRepository() repository.IAuthRepository
}