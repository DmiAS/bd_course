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
	return p.db.Updates(project).Error
}

func (p ProjectRepository) Get(id uuid.UUID) (*models.Project, error) {
	project := &models.Project{}
	res := p.db.Where("id = ?", id).First(project)
	return project, res.Error
}

func (p ProjectRepository) Delete(id uuid.UUID) error {
	return p.db.Where("id = ?", id).Delete(&models.Projects{}).Error
}

func (p ProjectRepository) GetAll() models.Projects {
	var projects models.Projects
	p.db.Find(&projects)
	return projects
}
