package project

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) Create(project *models.Project) error {
	ps := s.unit.GetProjectRepository()
	return ps.Create(project)
}

func (s Service) Get(clientID, projectID uuid.UUID) (*models.Project, error) {
	ps := s.unit.GetProjectRepository()
	return ps.Get(clientID, projectID)
}

func (s Service) GetAll(clientID uuid.UUID) *models.ProjectsList {
	ps := s.unit.GetProjectRepository()
	projects := ps.GetAll(clientID)
	return models.NewProjectsList(projects)
}

func (s Service) Update(project *models.Project) error {
	ps := s.unit.GetProjectRepository()
	return ps.Update(project)
}

func (s Service) Delete(clientID, projectID uuid.UUID) error {
	ps := s.unit.GetProjectRepository()
	return ps.Delete(clientID, projectID)
}