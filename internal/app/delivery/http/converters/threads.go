package converters

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertCreateThreadInput(req *ds.CreateThreadInput) *models.Thread {
	return &models.Thread{
		ProjectID: req.ProjectID,
		Name:      req.Name,
	}
}

func ConvertCreateThreadOutput(id uuid.UUID, name string) *ds.CreateThreadOutput {
	return &ds.CreateThreadOutput{
		ds.Thread{
			ID:   id,
			Name: name,
		},
	}
}

func ConvertGetThreadsOutput(projects models.Threads) *ds.GetThreadsOutput {
	cnt := len(projects)
	res := make([]ds.Thread, 0, cnt)

	for _, project := range projects {
		res = append(res, ds.Thread{
			ID:   project.ID,
			Name: project.Name,
		})
	}

	return &ds.GetThreadsOutput{
		Count:   cnt,
		Threads: res,
	}
}

func ConvertUpdateThreadInput(name string, id uuid.UUID) *models.Thread {
	return &models.Thread{
		ID:   id,
		Name: name,
	}
}
