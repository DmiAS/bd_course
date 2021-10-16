package worker

import (
	"log"

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
		return urep.Create(&models.Worker{
			ID:       authInfo.UserID,
			Grade:    worker.Grade,
			Position: worker.Position,
		})
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
		return wRep.Update(&models.Worker{
			ID:       worker.ID,
			Grade:    worker.Grade,
			Position: worker.Position,
		})
	})
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

	uRep := s.unit.GetUserRepository()
	user, err := uRep.Get(id)
	if err != nil {
		return nil, err
	}
	return &models.WorkerEntity{
		User:     *user,
		Grade:    worker.Grade,
		Position: worker.Position,
	}, nil
}

func (s *Service) GetAll() *models.WorkersList {
	wRep := s.unit.GetWorkerRepository()
	workers := wRep.GetAll()

	uRep := s.unit.GetUserRepository()
	users := uRep.GetAll(models.WorkerRole)
	length := len(workers)
	if len(users) < length {
		length = len(users)
		log.Printf("len users (%d) != workers (%d)\n", len(users), len(workers))
	}
	wks := make([]models.WorkerEntity, 0, length)
	for i := range workers {
		wks = append(wks, models.WorkerEntity{
			User:     users[i],
			Grade:    workers[i].Grade,
			Position: workers[i].Position,
		})
	}
	return &models.WorkersList{
		Amount:  len(workers),
		Workers: wks,
	}
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
