package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ThreadRepository struct {
	db *gorm.DB
}

func (t ThreadRepository) Create(thread *models.Thread) error {
	panic("implement me")
}

func (t ThreadRepository) Get(projectID, threadID uuid.UUID) (*models.Thread, error) {
	panic("implement me")
}

func (t ThreadRepository) GetAll(projectID uuid.UUID) models.Threads {
	panic("implement me")
}

func (t ThreadRepository) Update(thread *models.Thread) error {
	panic("implement me")
}

func (t ThreadRepository) Delete(projectID, threadID uuid.UUID) error {
	panic("implement me")
}

func NewThreadRepository(db *gorm.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}
