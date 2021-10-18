package orm

import (
	"log"
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	db *gorm.DB
}

func (c CampaignRepository) GetCampaign(campaignID uuid.UUID) (*models.Campaign, error) {
	camp := &models.Campaign{}
	res := c.db.Model(&models.Campaign{}).Where("id = ?", campaignID).First(camp)
	if err := res.Error; err != nil {
		return nil, err
	}
	return camp, nil
}

func (c CampaignRepository) GetCampaignStat(campID uuid.UUID, from, to time.Time) []models.CampaignStat {
	var stats []models.CampaignStat
	if err := c.db.
		Model(&models.CampaignStat{}).
		Where("camp_id = ? and date between symmetric ? and ?", campID, from, to).
		Find(&stats).Error; err != nil {
		log.Printf("[orm][stat] %s\n", err.Error())
	}
	return stats
}

func (c CampaignRepository) GetThreadCampaigns(threadID uuid.UUID, limit int, created int64) models.Campaigns {
	var camps models.Campaigns
	c.db.
		Limit(limit).
		Order("created desc").
		Where("created < ? and thread_id = ?", created, threadID).
		Find(&camps)
	return camps
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (c CampaignRepository) GetAll(limit int, created int64) models.Campaigns {
	var camps models.Campaigns
	c.db.
		Limit(limit).
		Order("created desc").
		Where("created < ?", created).
		Find(&camps)
	return camps
}

func (c CampaignRepository) GetCampaigns(workerID uuid.UUID, created int64, limit int) models.Campaigns {
	var camps models.Campaigns
	c.db.
		Limit(limit).
		Order("created desc").
		Where("targetolog_id = ? and created < ?", workerID, created).Find(&camps)
	return camps
}

func (c CampaignRepository) Update(camp *models.Campaign) error {
	return c.db.Where("id = ?", camp.ID).Updates(camp).Error
}
