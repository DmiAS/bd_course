package worker

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(worker *models.WorkerEntity) (*models.Auth, error) {
	worker.ID = uuid.UUID{}
	var authInfo *models.Auth
	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		aServ := auth.NewService(u)
		var err error
		authInfo, err = aServ.Create(worker.User.FirstName, worker.User.LastName, models.WorkerRole)
		if err != nil {
			return err
		}

		bdWorker := &models.Worker{User: worker.User, Grade: worker.Grade, Position: worker.Position}
		bdWorker.User.ID = authInfo.UserID
		wRep := u.GetWorkerRepository()
		if err := wRep.Create(bdWorker); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return authInfo, nil
}

func (s *Service) Update(worker *models.WorkerEntity) error {
	wRep := s.unit.GetWorkerRepository()
	bdWorker := &models.Worker{User: worker.User, Grade: worker.Grade, Position: worker.Position}
	return wRep.Update(bdWorker)
}

func (s *Service) Delete(id uuid.UUID) error {
	aRep := s.unit.GetAuthRepository()
	return aRep.Delete(id)
}

func (s *Service) Get(id uuid.UUID) (*models.WorkerEntity, error) {
	wRep := s.unit.GetWorkerRepository()
	worker, err := wRep.Get(id)
	if err != nil {
		return nil, err
	}
	bdWorker := &models.WorkerEntity{User: worker.User, Grade: worker.Grade, Position: worker.Position}
	return bdWorker, nil
}

func (s *Service) GetAll() *models.WorkersList {
	wRep := s.unit.GetWorkerRepository()
	workers := wRep.GetAll()
	wks := make([]models.WorkerEntity, 0, len(workers))
	for i := range workers {
		wks = append(wks, models.WorkerEntity{
			User:     workers[i].User,
			Grade:    workers[i].Grade,
			Position: workers[i].Position,
		})
	}
	return &models.WorkersList{
		Amount:  len(workers),
		Workers: wks,
	}
}
