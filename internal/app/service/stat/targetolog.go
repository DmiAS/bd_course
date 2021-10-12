package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) getThreadStat(threadID uuid.UUID, name string, from, to time.Time) models.ThreadSimpleStat {
	campsStat := s.getCampsStat(threadID, from, to)
	thread := collectThreadStat(campsStat)
	thread.ID = threadID
	thread.From = from
	thread.To = to
	thread.Name = name
	return thread
}

func (s Service) getCampsStat(threadID uuid.UUID, from, to time.Time) []models.CampSimpleStat {
	rep := s.unit.GetCampaignsRepository()
	camps := rep.GetThreadCampaigns(threadID)
	var total []models.CampSimpleStat
	for _, camp := range camps {
		stats := rep.GetCampaignStat(camp.ID, from, to)
		stat := collectCampaignStat(stats)
		stat.From = from
		stat.To = to
		stat.TargetologID = camp.TargetologID
		stat.CabinetID = camp.CabinetID
		stat.VkClientID = camp.ClientID
		stat.CampID = camp.ID
		total = append(total, stat)
	}
	return total
}

func collectThreadStat(stats []models.CampSimpleStat) models.ThreadSimpleStat {
	res := models.ThreadSimpleStat{}
	for _, stat := range stats {
		res.Sales += stat.Sales
		res.Spent += stat.Spent
		res.Unsubs += stat.Unsubs
		res.Subs += stat.Subs
		res.Spent += stat.Spent
		res.Conversion += stat.Conversion
		res.Impressions += stat.Impressions
	}
	return res
}

func collectCampaignStat(stats []models.CampaignStat) models.CampSimpleStat {
	res := models.CampSimpleStat{}
	for _, stat := range stats {
		res.Impressions += stat.Impressions
		res.Conversion += stat.Conversion
		res.Sales += stat.Sales
		res.Subs += len(stat.Subs)
		res.Unsubs += len(stat.Unsubs)
		res.Spent += stat.Spent
	}
	return res
}
