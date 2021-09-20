package thread

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/models"
)

func (s Service) Create(thread *models.Thread) (uuid.UUID, error) {
	return s.rep.Create(thread)
}

func (s Service) Get(projectID uuid.UUID) (models.Threads, error) {
	return s.rep.Get(projectID)
}

func (s Service) Update(thread *models.Thread) error {
	return s.rep.Update(thread)
}

func (s Service) Delete(id uuid.UUID) error {
	return s.rep.Delete(id)
}
