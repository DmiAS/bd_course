package thread

import (
	"github.com/DmiAS/bd_course/internal/app/repository"
)

type Service struct {
	rep repository.IThreadRepository
}

func NewService(rep repository.IThreadRepository) *Service {
	return &Service{rep: rep}
}
