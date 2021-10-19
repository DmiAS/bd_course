package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s Service) GetFullCampaignStat(campaignID, userID uuid.UUID, role models.Role, from, to time.Time) (*models.CampStat, error) {
	rep := s.unit.GetCampaignsRepository()
	camp, err := rep.GetCampaign(campaignID)
	if err != nil {
		return nil, err
	}
	// check access
	if role == models.ClientRole {
		return nil, errors.New("access denied")
	}
	if role == models.WorkerRole {
		if camp.TargetologID != userID {
			return nil, errors.New("access denied")
		}
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
		date := stat.Date.Format(models.TimeTemplate)
		days = append(days, models.CampaignDayStat{
			Day:         date,
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
