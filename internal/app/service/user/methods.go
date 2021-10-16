package user

import (
	"github.com/DmiAS/bd_course/internal/app/service/auth"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s *Service) Create(user *models.User) (*models.Auth, error) {
	var authInfo *models.Auth
	if err := s.unit.WithTransaction(func(u uwork.UnitOfWork) error {
		aServ := auth.NewService(u)
		var err error
		authInfo, err = aServ.Create(user.FirstName, user.LastName)
		if err != nil {
			return err
		}

		user.ID = authInfo.UserID
		cRep := u.GetUserRepository()
		if err := cRep.Create(user); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return authInfo, nil
}

func (s *Service) Update(user *models.User) error {
	cRep := s.unit.GetUserRepository()
	return cRep.Update(user)
}

func (s *Service) Get(id uuid.UUID) (*models.User, error) {
	cRep := s.unit.GetUserRepository()
	return cRep.Get(id)
}

func (s *Service) GetAll(role models.Role) *models.UserList {
	cRep := s.unit.GetUserRepository()
	users := cRep.GetAll(role)
	return &models.UserList{
		Amount: len(users),
		Users:  users,
	}
}
