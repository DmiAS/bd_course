package thread

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) Create(thread *models.Thread) error {
	rep := s.unit.GetThreadsRepository()
	return rep.Create(thread)
}

func (s Service) Get(threadID uuid.UUID) (*models.Thread, error) {
	rep := s.unit.GetThreadsRepository()
	return rep.Get(threadID)
}

func (s Service) GetAll(projectID uuid.UUID) *models.ThreadsList {
	rep := s.unit.GetThreadsRepository()
	threads := rep.GetAll(projectID)
	return models.NewThreadsList(threads)
}

func (s Service) Update(thread *models.Thread) error {
	rep := s.unit.GetThreadsRepository()
	return rep.Update(thread)
}

func (s Service) Delete(threadID uuid.UUID) error {
	rep := s.unit.GetThreadsRepository()
	return rep.Delete(threadID)
}
