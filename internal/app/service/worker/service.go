package worker

import (
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

type Service struct {
	unit uwork.UnitOfWork
}

func NewService(u uwork.UnitOfWork) *Service {
	return &Service{unit: u}
}
