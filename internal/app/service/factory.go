package service

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/service/campaign"
	"github.com/DmiAS/bd_course/internal/app/service/clients"
	"github.com/DmiAS/bd_course/internal/app/service/project"
	"github.com/DmiAS/bd_course/internal/app/service/thread"
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

type ClientFactory struct {
	unit uwork.UnitOfWork
}

func NewClientFactory(u uwork.UnitOfWork) *ClientFactory {
	return &ClientFactory{unit: u}
}
func (c ClientFactory) GetService(role uwork.Role) IClientService {
	unit := c.unit.WithRole(role)
	return clients.NewService(unit)
}

type ProjectFactory struct {
	unit uwork.UnitOfWork
}

func NewProjectFactory(u uwork.UnitOfWork) *ProjectFactory {
	return &ProjectFactory{unit: u}
}
func (p ProjectFactory) GetService(role uwork.Role) IProjectService {
	unit := p.unit.WithRole(role)
	return project.NewService(unit)
}

type ThreadFactory struct {
	unit uwork.UnitOfWork
}

func NewThreadFactory(u uwork.UnitOfWork) *ThreadFactory {
	return &ThreadFactory{unit: u}
}
func (t ThreadFactory) GetService(role uwork.Role) IThreadService {
	unit := t.unit.WithRole(role)
	return thread.NewService(unit)
}

type CampaignFactory struct {
	unit uwork.UnitOfWork
}

func NewCampaignsFactory(u uwork.UnitOfWork) *CampaignFactory {
	return &CampaignFactory{unit: u}
}
func (c CampaignFactory) GetService(role uwork.Role) ICampaignService {
	unit := c.unit.WithRole(role)
	return campaign.NewService(unit)
}
