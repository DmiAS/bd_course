package worker

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(worker *models.Worker) (*models.Auth, error) {
	var authInfo *models.Auth
	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		aServ := auth.NewService(u)
		var err error
		authInfo, err = aServ.Create(worker.User.FirstName, worker.User.LastName, models.WorkerRole)
		if err != nil {
			return err
		}

		worker.User.ID = authInfo.UserID
		wRep := u.GetWorkerRepository()
		if err := wRep.Create(worker); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return authInfo, nil
}

func (s *Service) Update(worker *models.Worker) error {
	wRep := s.unit.GetWorkerRepository()
	return wRep.Update(worker)
}

func (s *Service) Delete(id uuid.UUID) error {
	aRep := s.unit.GetAuthRepository()
	return aRep.Delete(id)
}

func (s *Service) Get(id uuid.UUID) (*models.Worker, error) {
	wRep := s.unit.GetWorkerRepository()
	return wRep.Get(id)
}

func (s *Service) GetAll() models.Workers {
	wRep := s.unit.GetWorkerRepository()
	return wRep.GetAll()
}
