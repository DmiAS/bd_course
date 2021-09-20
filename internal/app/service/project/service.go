package project

import "github.com/DmiAS/bd_course/internal/app/repository"

type Service struct {
	rep repository.IProjectRepository
}

func NewService(rep repository.IProjectRepository) *Service {
	return &Service{rep: rep}
}
