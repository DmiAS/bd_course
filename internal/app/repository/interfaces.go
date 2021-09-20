package repository

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

type IWorkerRepository interface {
	CreateWorker(worker *models.Worker) (uuid.UUID, error)
}

type IAuthRepository interface {
	CreateAuth(auth *models.Auth) error
}
