package project

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) Create(project *models.Project) (uuid.UUID, error) {
	panic("implement me")
}

func (s Service) Get(clientID uuid.UUID) (*models.Project, error) {
	panic("implement me")
}

func (s Service) GetAll() models.Projects {
	panic("implement me")
}

func (s Service) Update(project *models.Project) error {
	panic("implement me")
}

func (s Service) Delete(id uuid.UUID) error {
	panic("implement me")
}
