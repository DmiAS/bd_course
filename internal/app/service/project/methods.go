package project

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s Service) Create(clientID uuid.UUID, name string) error {
	ps := s.unit.GetProjectRepository()
	project := &models.Project{
		ID:       uuid.New(),
		ClientID: clientID,
		Name:     name,
	}
	return ps.Create(project)
}

func (s Service) Get(projectID, userID uuid.UUID, role models.Role) (*models.Project, error) {
	ps := s.unit.GetProjectRepository()
	project, err := ps.Get(projectID)
	if err != nil {
		return nil, err
	}
	// check access
	if role == models.ClientRole {
		if project.ClientID != userID {
			return nil, errors.New("access denied")
		}
	}
	return project, nil
}

func (s Service) GetAll(clientID uuid.UUID, pagination *models.Pagination) *models.ProjectsList {
	ps := s.unit.GetProjectRepository()
	pag := models.GetPaginationInfo(pagination)
	projects := ps.GetAll(clientID, pag.Cursor, pag.Limit)
	return createProjectList(projects)
}

func (s Service) Update(project *models.Project) error {
	ps := s.unit.GetProjectRepository()
	return ps.Update(project)
}

func (s Service) Delete(projectID uuid.UUID) error {
	ps := s.unit.GetProjectRepository()
	return ps.Delete(projectID)
}

func createProjectList(projects models.Projects) *models.ProjectsList {
	var cursor int64
	if len(projects)-1 >= 0 {
		cursor = projects[len(projects)-1].Created
	}
	return &models.ProjectsList{
		Cursor:   cursor,
		Projects: projects,
		Amount:   len(projects),
	}
}
