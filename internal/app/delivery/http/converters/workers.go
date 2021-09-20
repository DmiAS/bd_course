package converters

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertWorkerInput(req *ds.CreateWorkerInput) (*models.Worker, string) {
	return &models.Worker{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Grade:     req.Grade,
		Position:  req.Position,
	}, req.Login
}

func ConvertWorkeroutput(login, password string) *ds.CreateWorkerOutput {
	return &ds.CreateWorkerOutput{
		Login:    login,
		Password: password,
	}
}
