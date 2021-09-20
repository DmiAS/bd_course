package worker

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) CreateWorker(worker *models.Worker) (uuid.UUID, error) {
	id, err := s.rep.CreateWorker(worker)
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}
