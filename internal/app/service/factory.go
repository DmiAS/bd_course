package service

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/service/worker"
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

type WorkerFactory struct {
	unit uwork.UnitOfWork
}

func NewWorkerFactory(u uwork.UnitOfWork) *WorkerFactory {
	return &WorkerFactory{unit: u}
}
func (w WorkerFactory) GetService(role uwork.Role) IWorkerService {
	unit := w.unit.WithRole(role)
	return worker.NewService(unit)
}

type AuthFactory struct {
	unit uwork.UnitOfWork
}

func NewAuthFactory(u uwork.UnitOfWork) *AuthFactory {
	return &AuthFactory{unit: u}
}
func (a AuthFactory) GetService(role uwork.Role) IAuthService {
	unit := a.unit.WithRole(role)
	return auth.NewService(unit)
}
