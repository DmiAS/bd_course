package stat

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) GetProjectStat(projectID uuid.UUID) (*models.ProjectStat, error) {
	panic("implement me")
}

func (s Service) GetThreadStat(threadID uuid.UUID) (*models.ThreadStat, error) {
	panic("implement me")
}

func (s Service) GetTargetologStat(targetologID uuid.UUID) (*models.TargetologStat, error) {
	panic("implement me")
}
