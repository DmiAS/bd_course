package converters

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertCreateProjectInput(req *ds.CreateProjectInput) *models.Project {
	return &models.Project{
		ID:       uuid.UUID{},
		ClientID: req.ClientID,
		Name:     req.Name,
	}
}

func ConvertCreateProjectOutput(id uuid.UUID, name string) *ds.CreateProjectOutput {
	return &ds.CreateProjectOutput{
		ds.Project{
			ID:   id,
			Name: name,
		},
	}
}

func ConvertGetProjectsOutput(projects models.Projects) *ds.GetProjectsOutput {
	cnt := len(projects)
	res := make([]ds.Project, 0, cnt)

	for _, project := range projects {
		res = append(res, ds.Project{
			ID:   project.ID,
			Name: project.Name,
		})
	}

	return &ds.GetProjectsOutput{
		Count:    cnt,
		Projects: res,
	}
}

func ConvertUpdateInput(name string, id uuid.UUID) *models.Project {
	return &models.Project{
		ID:   id,
		Name: name,
	}
}
