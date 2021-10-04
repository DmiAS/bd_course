package thread

import (
	"github.com/DmiAS/bd_course/internal/app/repository"
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

type Service struct {
	unit uwork.UnitOfWork
	rep  repository.IThreadRepository
}

func NewService(rep repository.IThreadRepository) *Service {
	return &Service{rep: rep}
}
