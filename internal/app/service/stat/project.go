package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) GetProjectStat(projectID uuid.UUID, from, to time.Time) (*models.ProjectStat, error) {
	rep := s.unit.GetProjectRepository()
	project, err := rep.Get(projectID)
	if err != nil {
		return nil, err
	}
	threadsStat := s.getThreadsStat(projectID, from, to)
	projectStat := collectProjectStat(threadsStat)
	projectStat.ProjectID = projectID
	projectStat.From = from
	projectStat.To = to
	projectStat.Threads = threadsStat
	projectStat.Name = project.Name
	return &projectStat, nil
}

func (s Service) getThreadsStat(projectID uuid.UUID, from, to time.Time) []models.ThreadSimpleStat {
	rep := s.unit.GetThreadsRepository()
	threads := rep.GetAll(projectID)
	var stats []models.ThreadSimpleStat
	for _, thread := range threads {
		stats = append(stats, s.getThreadSimpleStat(thread, from, to))
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
