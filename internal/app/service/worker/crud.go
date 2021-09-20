package worker

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(worker *models.Worker) (uuid.UUID, error) {
	id, err := s.rep.Create(worker)
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func (s *Service) Update(worker *models.Worker) error {
	return s.rep.Update(worker)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.rep.Delete(id)
}

func (s *Service) Get(id uuid.UUID) (*models.Worker, error) {
	return s.rep.Get(id)
}

func (s *Service) GetAll() (models.Workers, error) {
	return s.rep.GetAll()
}
