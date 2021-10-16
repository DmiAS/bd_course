package service

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/service/campaign"
	"github.com/DmiAS/bd_course/internal/app/service/project"
	"github.com/DmiAS/bd_course/internal/app/service/stat"
	"github.com/DmiAS/bd_course/internal/app/service/thread"
	"github.com/DmiAS/bd_course/internal/app/service/user"
	"github.com/DmiAS/bd_course/internal/app/service/worker"
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

type WorkerFactory struct {
	unit uwork.UnitOfWork
}

func NewWorkerFactory(u uwork.UnitOfWork) *WorkerFactory {
	return &WorkerFactory{unit: u}
}
func (w WorkerFactory) GetService(role models.Role) IWorkerService {
	unit := w.unit.WithRole(role)
	return worker.NewService(unit)
}

type AuthFactory struct {
	unit uwork.UnitOfWork
}

func NewAuthFactory(u uwork.UnitOfWork) *AuthFactory {
	return &AuthFactory{unit: u}
}
func (a AuthFactory) GetService(role models.Role) IAuthService {
	unit := a.unit.WithRole(role)
	return auth.NewService(unit)
}

type UserFactory struct {
	unit uwork.UnitOfWork
}

func NewUserFactory(u uwork.UnitOfWork) *UserFactory {
	return &UserFactory{unit: u}
}
func (u UserFactory) GetService(role models.Role) IUserService {
	unit := u.unit.WithRole(role)
	return user.NewService(unit)
}

type ProjectFactory struct {
	unit uwork.UnitOfWork
}

func NewProjectFactory(u uwork.UnitOfWork) *ProjectFactory {
	return &ProjectFactory{unit: u}
}
func (p ProjectFactory) GetService(role models.Role) IProjectService {
	unit := p.unit.WithRole(role)
	return project.NewService(unit)
}

type ThreadFactory struct {
	unit uwork.UnitOfWork
}

func NewThreadFactory(u uwork.UnitOfWork) *ThreadFactory {
	return &ThreadFactory{unit: u}
}
func (t ThreadFactory) GetService(role models.Role) IThreadService {
	unit := t.unit.WithRole(role)
	return thread.NewService(unit)
}

type CampaignFactory struct {
	unit uwork.UnitOfWork
}

func NewCampaignsFactory(u uwork.UnitOfWork) *CampaignFactory {
	return &CampaignFactory{unit: u}
}
func (c CampaignFactory) GetService(role models.Role) ICampaignService {
	unit := c.unit.WithRole(role)
	return campaign.NewService(unit)
}

type StatsFactory struct {
	unit uwork.UnitOfWork
}

func NewStatsFactory(u uwork.UnitOfWork) *StatsFactory {
	return &StatsFactory{unit: u}
}
func (s StatsFactory) GetService(role models.Role) IStatService {
	unit := s.unit.WithRole(role)
	return stat.NewService(unit)
}
