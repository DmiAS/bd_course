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

func (w WorkerRepository) Get(id uuid.UUID) (*models.Worker, error) {
	worker := &models.Worker{}
	res := w.db.Where("id = ?", id).First(worker)
	return worker, res.Error
}

func (w WorkerRepository) GetAll() models.Workers {
	var workers models.Workers
	w.db.Find(&workers)
	return workers
}

func (w WorkerRepository) GetCampaigns(workerID uuid.UUID) models.Campaigns {
	var camps models.Campaigns
	w.db.Model(&models.Campaign{}).Where("targetolog_id = ?", workerID).Find(&camps)
	return camps
}

func (w WorkerRepository) AttachCampaign(threadID, campID uuid.UUID) error {
	return w.db.Model(&models.Campaign{}).
		Where("id = ?", campID).
		Update("thread_id", threadID).Error
}

func (w WorkerRepository) AssignCampaign(camp *models.Campaign) error {
	return w.db.Create(camp).Error
}

func (w WorkerRepository) UnAssignCampaign(campID uuid.UUID) error {
	return w.db.Model(&models.Campaign{}).
		Where("id = ?", campID).
		UpdateColumn("targetolog_id", uuid.UUID{}).Error
}
