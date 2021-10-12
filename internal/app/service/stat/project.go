package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) GetProjectStat(projectID uuid.UUID, from, to time.Time) models.ProjectStat {
	threadsStat := s.getThreadsStat(projectID, from, to)
	project := collectProjectStat(threadsStat)
	project.ProjectID = projectID
	project.From = from
	project.To = to
	project.Threads = threadsStat
	return project
}

func (s Service) getThreadsStat(projectID uuid.UUID, from, to time.Time) []models.ThreadSimpleStat {
	rep := s.unit.GetThreadsRepository()
	threads := rep.GetAll(projectID)
	var stats []models.ThreadSimpleStat
	for _, thread := range threads {
		stats = append(stats, s.getThreadStat(thread.ID, thread.Name, from, to))
	}
	return stats
}

func collectProjectStat(stats []models.ThreadSimpleStat) models.ProjectStat {
	res := models.ProjectStat{}
	for _, stat := range stats {
		res.Impressions += stat.Impressions
		res.Spent += stat.Spent
		res.Conversion += stat.Conversion
		res.Sales += stat.Sales
		res.Subs += stat.Subs
		res.Unsubs += stat.Unsubs
	}
	return res
}
