package auth

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/uwork"
)

type Service struct {
	unit           uwork.UnitOfWork
	expireDuration time.Duration
	signKey        []byte
}

const (
	defaultExpireDuration = time.Minute * 30
	signKey               = "key"
)

func NewService(unit uwork.UnitOfWork) *Service {
	return &Service{
		unit:           unit,
		expireDuration: defaultExpireDuration,
		signKey:        []byte(signKey),
	}
}
