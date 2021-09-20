package worker

import "github.com/DmiAS/bd_course/internal/app/models"

func (s *Service) UpdateWorker(worker *models.Worker) error {
	return s.rep.Update(worker)
}
