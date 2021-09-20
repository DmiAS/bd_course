package converters

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertWorkerCreateInput(req *ds.CreateWorkerInput) (*models.Worker, string) {
	return &models.Worker{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Grade:     req.Grade,
		Position:  req.Position,
	}, req.Login
}

func ConvertWorkerCreateOutput(login, password string) *ds.CreateWorkerOutput {
	return &ds.CreateWorkerOutput{
		Login:    login,
		Password: password,
	}
}

func ConvertWorkerUpdateInput(worker *ds.UpdateWorkerInput, id uuid.UUID) *models.Worker {
	return &models.Worker{
		UUID:      id,
		FirstName: worker.FirstName,
		LastName:  worker.LastName,
		Grade:     worker.Grade,
		Position:  worker.Position,
	}
}

func convertToWorker(worker models.Worker) ds.Worker {
	return ds.Worker{
		FirstName: worker.FirstName,
		LastName:  worker.LastName,
		Grade:     worker.Grade,
		Position:  worker.Position,
	}
}

func ConvertWorkerGetOutput(worker *models.Worker) *ds.GetWorkerOutput {
	return &ds.GetWorkerOutput{convertToWorker(*worker)}
}

func ConvertWorkerGetAllOutput(workers models.Workers) *ds.GetAllWorkersOutput {
	cnt := len(workers)
	ws := make([]ds.WorkerUUID, 0, cnt)

	for _, w := range workers {
		ws = append(ws, ds.WorkerUUID{
			Worker: convertToWorker(w),
			UUID:   w.UUID,
		})
	}

	return &ds.GetAllWorkersOutput{
		Count:   cnt,
		Workers: ws,
	}
}
