package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s Service) GetProjectStat(projectID, userID uuid.UUID, role models.Role, from, to time.Time) (*models.ProjectStat, error) {
	rep := s.unit.GetProjectRepository()
	project, err := rep.Get(projectID)
	if err != nil {
		return nil, err
	}
	// check access
	if (role == models.ClientRole && project.ClientID != userID) || role == models.WorkerRole {
		return nil, errors.New("access denied")
	}
	threadsStat := s.getThreadsStat(projectID, from, to)
	projectStat := collectProjectStat(threadsStat)
	projectStat.ProjectID = projectID
	projectStat.From = from.Format(models.TimeTemplate)
	projectStat.To = to.Format(models.TimeTemplate)
	projectStat.Threads = threadsStat
	projectStat.Name = project.Name
	return &projectStat, nil
}

func (s Service) getThreadsStat(projectID uuid.UUID, from, to time.Time) []models.ThreadSimpleStat {
	rep := s.unit.GetThreadsRepository()
	threads := rep.GetAll(projectID, 0, 0)
	var stats []models.ThreadSimpleStat
	for _, thread := range threads {
		threadStat := s.getThreadSimpleStat(thread, from, to)
		if threadStat.Spent > models.Eps {
			stats = append(stats, threadStat)
		}
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
