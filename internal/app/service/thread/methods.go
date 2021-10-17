package thread

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
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

func (s Service) Get(threadID uuid.UUID) (*models.Thread, error) {
	rep := s.unit.GetThreadsRepository()
	return rep.Get(threadID)
}

func (s Service) GetAll(projectID uuid.UUID, pagination *models.Pagination) *models.ThreadsList {
	rep := s.unit.GetThreadsRepository()
	pag := models.GetPaginationInfo(pagination)
	threads := rep.GetAll(projectID, pag.Cursor, pag.Limit)
	return createThreadList(threads)
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
