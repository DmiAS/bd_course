package worker

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(worker *models.Worker) (*models.Auth, error) {
	unit := s.unit.WithTransaction()
	aRep := auth.NewService(unit)
	authInfo, err := aRep.CreateAuth(worker.FirstName, worker.LastName, models.WorkerRole)
	if err != nil {
		unit.Rollback()
		return nil, err
	}

	worker.ID = authInfo.ID
	wRep := unit.GetWorkerRepository()
	if err := wRep.Create(worker); err != nil {
		unit.Rollback()
		return nil, err
	}

	unit.Commit()
	return authInfo, nil
}

func (s *Service) Update(worker *models.Worker) error {
	wRep := s.unit.GetWorkerRepository()
	return wRep.Update(worker)
}

func (s *Service) Delete(id uuid.UUID) error {
	wRep := s.unit.GetWorkerRepository()
	return wRep.Delete(id)
}

func (s *Service) Get(id uuid.UUID) (*models.Worker, error) {
	wRep := s.unit.GetWorkerRepository()
	return wRep.Get(id)
}

func (s *Service) GetAll() (models.Workers, error) {
	wRep := s.unit.GetWorkerRepository()
	return wRep.GetAll()
}
