package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkerRepository struct {
	db *gorm.DB
}

func NewWorkerRepository(db *gorm.DB) *WorkerRepository {
	return &WorkerRepository{db: db}
}

func (w WorkerRepository) Create(worker *models.Worker) error {
	return w.db.Create(worker).Error
}

func (w WorkerRepository) Update(worker *models.Worker) error {
	return w.db.Updates(worker).Error
}

func (w WorkerRepository) Get(id uuid.UUID) (*models.WorkerEntity, error) {
	worker := &models.WorkerEntity{}
	if err := w.db.
		Model(&models.Worker{}).
		Select("*").
		Joins("join users on workers.user_id = users.id and users.id = ?", id).
		First(worker).
		Error; err != nil {
		return nil, err
	}
	return worker, nil
}

func (w WorkerRepository) GetAll() models.Workers {
	var workers models.Workers
	w.db.Find(&workers)
	return workers
}
