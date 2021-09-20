package converters

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertWorkerInput(req *ds.Worker) *models.Worker {
	return &models.Worker{
		User: models.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			VkLink:    req.VkLink,
			TgLink:    req.TgLink,
		},
	}
}

func ConvertUpdateWorkerInput(req *ds.Worker, id uuid.UUID) *models.Worker {
	return &models.Worker{
		User: models.User{
			ID:        id,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			VkLink:    req.VkLink,
			TgLink:    req.TgLink,
		},
	}
}

func convertToWorker(worker models.Worker) ds.Worker {
	return ds.Worker{
		User:     convertUser(worker.User),
		Grade:    worker.Grade,
		Position: worker.Position,
	}
}

func ConvertWorkerOutput(worker *models.Worker) *ds.Worker {
	res := convertToWorker(*worker)
	return &res
}

func ConvertGetAllWorkerOutput(workers models.Workers) *ds.GetAllWorkersOutput {
	cnt := len(workers)
	ws := make([]ds.WorkerUUID, 0, cnt)

	for _, w := range workers {
		ws = append(ws, ds.WorkerUUID{
			Worker: convertToWorker(w),
			UUID:   w.User.ID,
		})
	}

	return &ds.GetAllWorkersOutput{
		Count:   cnt,
		Workers: ws,
	}
}
