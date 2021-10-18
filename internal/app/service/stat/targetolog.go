package stat

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) GetTargetologStat(targetologID uuid.UUID, from, to time.Time) (*models.TargetologStat, error) {
	rep := s.unit.GetCampaignsRepository()
	camps := rep.GetTargetologCampaigns(targetologID, 0, 0)
	campsStat := s.getCampsStat(camps, from, to)
	return &models.TargetologStat{Camps: campsStat}, nil
}
