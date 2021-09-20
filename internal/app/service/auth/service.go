package auth

import (
	"github.com/DmiAS/bd_course/internal/app/repository"
)

type Service struct {
	rep repository.IAuthRepository
}

func NewService(rep repository.IAuthRepository) *Service {
	return &Service{rep: rep}
}
