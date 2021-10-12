package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (c CampaignRepository) GetAll() models.Campaigns {
	var camps models.Campaigns
	c.db.Find(&camps)
	return camps
}

func (c CampaignRepository) GetCampaigns(workerID uuid.UUID) models.Campaigns {
	var camps models.Campaigns
	c.db.Model(&models.Campaign{}).Where("targetolog_id = ?", workerID).Find(&camps)
	return camps
}

func (c CampaignRepository) Update(camp *models.Campaign) error {
	return c.db.Where("id = ?", camp.ID).Updates(camp).Error
}
