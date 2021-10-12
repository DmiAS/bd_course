package stat

import (
	"github.com/DmiAS/bd_course/internal/app/uwork"
)

type Service struct {
	unit uwork.UnitOfWork
}

func NewService(unit uwork.UnitOfWork) *Service {
	return &Service{unit: unit}
}
