package worker

import (
	"github.com/DmiAS/bd_course/internal/app/service/user"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(worker *models.WorkerEntity) (*models.Auth, error) {
	worker.ID = uuid.UUID{}
	var authInfo *models.Auth
	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		us := user.NewService(u)
		var err error
		authInfo, err = us.Create(createUser(worker))
		if err != nil {
			return err
		}

		urep := u.GetWorkerRepository()
		return urep.Create(createWorker(authInfo.UserID, worker))
	}); err != nil {
		return nil, err
	}
	return authInfo, nil
}

func (s *Service) Update(worker *models.WorkerEntity) error {
	return s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		uRep := s.unit.GetUserRepository()
		if err := uRep.Update(createUser(worker)); err != nil {
			return err
		}

		wRep := s.unit.GetWorkerRepository()
		return wRep.Update(createWorker(worker.ID, worker))
	})
}

func (s *Service) Delete(id uuid.UUID) error {
	aRep := s.unit.GetAuthRepository()
	return aRep.Delete(id)
}

func (s *Service) Get(id uuid.UUID) (*models.WorkerEntity, error) {
	wRep := s.unit.GetWorkerRepository()
	worker, err := wRep.Get(id)
	return worker, err
}

func (s *Service) GetAll(pagination *models.Pagination) *models.WorkersList {
	wRep := s.unit.GetWorkerRepository()
	pag := models.GetPaginationInfo(pagination)
	workers := wRep.GetAll(pag.Cursor, pag.Limit)

	return createWorkerList(workers)
}

func createUser(worker *models.WorkerEntity) *models.User {
	return &models.User{
		ID:        worker.ID,
		FirstName: worker.FirstName,
		LastName:  worker.LastName,
		VkLink:    worker.VkLink,
		TgLink:    worker.TgLink,
		Role:      models.WorkerRole,
	}
}

func createWorker(id uuid.UUID, wk *models.WorkerEntity) *models.Worker {
	return &models.Worker{
		UserID:   id,
		Grade:    wk.Grade,
		Position: wk.Position,
	}
}

func createWorkerList(workers []models.WorkerEntity) *models.WorkersList {
	var cursor int64
	if len(workers)-1 >= 0 {
		cursor = workers[len(workers)-1].Created
	}
	return &models.WorkersList{
		Cursor:  cursor,
		Amount:  len(workers),
		Workers: workers,
	}
}
