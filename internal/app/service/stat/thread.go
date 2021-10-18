package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) GetThreadStat(threadID uuid.UUID, from, to time.Time) (*models.ThreadStat, error) {
	rep := s.unit.GetThreadsRepository()
	thread, err := rep.Get(threadID)
	if err != nil {
		return nil, err
	}
	camps := s.getThreadCampaigns(threadID)
	campsStat := s.getCampsStat(camps, from, to)
	targets := collectThreadTargetologs(camps)
	return &models.ThreadStat{
		ThreadSimpleStat: createThreadSimpleStat(campsStat, *thread, from, to),
		Targetologs:      targets,
		Camps:            campsStat,
	}, nil
}

func (s Service) getThreadSimpleStat(thread models.Thread, from, to time.Time) models.ThreadSimpleStat {
	camps := s.getThreadCampaigns(thread.ID)
	campsStat := s.getCampsStat(camps, from, to)
	stat := createThreadSimpleStat(campsStat, thread, from, to)
	return stat
}

func (s Service) getThreadCampaigns(threadID uuid.UUID) models.Campaigns {
	rep := s.unit.GetCampaignsRepository()
	return rep.GetThreadCampaigns(threadID, 0, 0)
}

func (s Service) getCampsStat(camps models.Campaigns, from, to time.Time) []models.CampSimpleStat {
	rep := s.unit.GetCampaignsRepository()
	var total []models.CampSimpleStat
	for _, camp := range camps {
		stats := rep.GetCampaignStat(camp.ID, from, to)
		stat := collectCampaignStat(stats)
		if stat.Spent < models.Eps {
			continue
		}
		stat.From = from.Format(models.TimeTemplate)
		stat.To = to.Format(models.TimeTemplate)
		stat.TargetologID = camp.TargetologID
		stat.CabinetID = camp.CabinetID
		stat.VkClientID = camp.ClientID
		stat.CampID = camp.ID
		stat.Name = camp.Name
		total = append(total, stat)
	}
	return total
}

func createThreadSimpleStat(campsStat []models.CampSimpleStat, thread models.Thread, from, to time.Time) models.ThreadSimpleStat {
	stat := collectThreadStat(campsStat)
	stat.ID = thread.ID
	stat.From = from.Format(models.TimeTemplate)
	stat.To = to.Format(models.TimeTemplate)
	stat.Name = thread.Name
	return stat
}

func collectThreadTargetologs(camps models.Campaigns) []uuid.UUID {
	m := make(map[uuid.UUID]struct{})
	for _, camp := range camps {
		if _, ok := m[camp.TargetologID]; !ok {
			m[camp.TargetologID] = struct{}{}
		}
	}
	ids := make([]uuid.UUID, 0, len(m))
	for key := range m {
		ids = append(ids, key)
	}
	return ids
}

func collectThreadStat(stats []models.CampSimpleStat) models.ThreadSimpleStat {
	res := models.ThreadSimpleStat{}
	for _, stat := range stats {
		res.Sales += stat.Sales
		res.Spent += stat.Spent
		res.Unsubs += stat.Unsubs
		res.Subs += stat.Subs
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
