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
	return t.db.Create(thread).Error
}

func (t ThreadRepository) Get(threadID uuid.UUID) (*models.Thread, error) {
	thread := &models.Thread{}
	res := t.db.
		Where("id = ?", threadID).First(thread)
	return thread, res.Error
}

func (t ThreadRepository) GetAll(projectID uuid.UUID) models.Threads {
	var threads models.Threads
	t.db.Where("project_id = ?", projectID).Find(&threads)
	return threads
}

func (t ThreadRepository) Update(thread *models.Thread) error {
	return t.db.
		Where("id = ?", thread.ID).Updates(thread).Error
}

func (t ThreadRepository) Delete(threadID uuid.UUID) error {
	return t.db.Where("id = ?", threadID).Delete(&models.Projects{}).Error
}

func NewThreadRepository(db *gorm.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}
