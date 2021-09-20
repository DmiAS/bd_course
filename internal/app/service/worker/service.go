package worker

import (
	"github.com/DmiAS/bd_course/internal/app/repository"
)

type Service struct {
	rep repository.IWorkerRepository
}

func NewService(rep repository.IWorkerRepository) *Service {
	return &Service{rep: rep}
}
