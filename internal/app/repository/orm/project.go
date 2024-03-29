package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (p ProjectRepository) Create(project *models.Project) error {
	return p.db.Create(project).Error
}

func (p ProjectRepository) Update(project *models.Project) error {
	return p.db.
		Where("id = ?", project.ID).Updates(project).Error
}

func (p ProjectRepository) Get(projectID uuid.UUID) (*models.Project, error) {
	project := &models.Project{}
	res := p.db.Where("id = ?", projectID).First(project)
	return project, res.Error
}

func (p ProjectRepository) Delete(projectID uuid.UUID) error {
	return p.db.Where("id = ?", projectID).Delete(&models.Projects{}).Error
}

func (p ProjectRepository) GetAll(clientID uuid.UUID, created int64, limit int) models.Projects {
	var projects models.Projects
	if created == 0 && limit == 0 {
		p.db.
			Where("client_id = ?", clientID).
			Find(&projects)
	} else {
		p.db.
			Limit(limit).
			Order("created desc").
			Where("client_id = ? and created < ?", clientID, created).
			Find(&projects)
	}
	return projects
}
