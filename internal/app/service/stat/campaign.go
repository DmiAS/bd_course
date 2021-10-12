package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) GetFullCampaignStat(campaignID uuid.UUID, from, to time.Time) (*models.CampStat, error) {
	rep := s.unit.GetCampaignsRepository()
	camp, err := rep.GetCampaign(campaignID)
	if err != nil {
		return nil, err
	}
	days := s.getCampDaysStat(campaignID, from, to)
	return &models.CampStat{
		CampID:     campaignID,
		CabinetID:  camp.CabinetID,
		VkClientID: camp.ClientID,
		Name:       camp.Name,
		Days:       days,
	}, nil
}

func (s Service) getCampDaysStat(campaignID uuid.UUID, from, to time.Time) []models.CampaignDayStat {
	rep := s.unit.GetCampaignsRepository()
	stats := rep.GetCampaignStat(campaignID, from, to)
	days := make([]models.CampaignDayStat, 0, len(stats))
	for _, stat := range stats {
		days = append(days, models.CampaignDayStat{
			Day:         stat.Date,
			Spent:       stat.Spent,
			Impressions: stat.Impressions,
			Conversion:  stat.Conversion,
			Subs:        len(stat.Subs),
			Unsubs:      len(stat.Unsubs),
			Sales:       stat.Sales,
		})
	}
	return days
}
