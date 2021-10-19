package thread

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s Service) Create(projectID uuid.UUID, name string) error {
	rep := s.unit.GetThreadsRepository()
	thread := &models.Thread{
		ID:        uuid.New(),
		ProjectID: projectID,
		Name:      name,
	}
	return rep.Create(thread)
}

func (s Service) Get(threadID uuid.UUID, userID uuid.UUID, role models.Role) (*models.Thread, error) {
	rep := s.unit.GetThreadsRepository()
	thread, err := rep.Get(threadID)
	if err != nil {
		return nil, err
	}
	// check access to this thread
	if role == models.ClientRole {
		rep := s.unit.GetProjectRepository()
		projects := rep.GetAll(userID, 0, 0)
		accessGranted := false
		for i := range projects {
			if projects[i].ID == thread.ProjectID {
				accessGranted = true
				break
			}
		}
		if !accessGranted {
			return nil, errors.New("access denied")
		}
	}
	return thread, nil
}

func (s Service) GetAll(projectID uuid.UUID, userID uuid.UUID, role models.Role, pagination *models.Pagination) (*models.ThreadsList, error) {
	// check access to this project
	if role == models.ClientRole {
		rep := s.unit.GetProjectRepository()
		project, err := rep.Get(projectID)
		if err != nil {
			return nil, err
		}
		if project.ClientID != userID {
			return nil, errors.New("access denied")
		}
	}
	rep := s.unit.GetThreadsRepository()
	pag := models.GetPaginationInfo(pagination)
	threads := rep.GetAll(projectID, pag.Cursor, pag.Limit)
	return createThreadList(threads), nil
}

func (s Service) Update(thread *models.Thread) error {
	rep := s.unit.GetThreadsRepository()
	return rep.Update(thread)
}

func (s Service) Delete(threadID uuid.UUID) error {
	rep := s.unit.GetThreadsRepository()
	return rep.Delete(threadID)
}

func createThreadList(threads models.Threads) *models.ThreadsList {
	var cursor int64
	if len(threads)-1 >= 0 {
		cursor = threads[len(threads)-1].Created
	}
	return &models.ThreadsList{
		Cursor:  cursor,
		Threads: threads,
		Amount:  len(threads),
	}
}
