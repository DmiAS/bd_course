package project

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(project *models.Project) (uuid.UUID, error) {
	return s.rep.Create(project)
}

func (s *Service) Get(clientID uuid.UUID) (models.Projects, error) {
	return s.rep.Get(clientID)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.rep.Delete(id)
}
