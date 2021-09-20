package auth

import "github.com/DmiAS/bd_course/internal/app/repository"

type Service struct {
	rep repository.IAuthRepository
}

func NewWorkerService(rep repository.IAuthRepository) *Service {
	return &Service{rep: rep}
}
