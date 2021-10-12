package clients

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(client *models.Client) (*models.Auth, error) {
	var authInfo *models.Auth
	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		aServ := auth.NewService(u)
		var err error
		authInfo, err = aServ.Create(client.User.FirstName, client.User.LastName, models.WorkerRole)
		if err != nil {
			return err
		}

		client.User.ID = authInfo.UserID
		cRep := u.GetClientRepository()
		if err := cRep.Create(client); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return authInfo, nil
}

func (s *Service) Update(client *models.Client) error {
	cRep := s.unit.GetClientRepository()
	return cRep.Update(client)
}

func (s *Service) Get(id uuid.UUID) (*models.Client, error) {
	cRep := s.unit.GetClientRepository()
	return cRep.Get(id)
}

func (s *Service) GetAll() *models.ClientsList {
	cRep := s.unit.GetClientRepository()
	clients := cRep.GetAll()
	return models.NewClientsList(clients)
}
