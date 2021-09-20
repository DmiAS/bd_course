package repository

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerRepository interface {
	Create(worker *models.Worker) (uuid.UUID, error)
	Update(worker *models.Worker) error
}

type IAuthRepository interface {
	Create(auth *models.Auth) error
}
