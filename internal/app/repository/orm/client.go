package orm

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c ClientRepository) Create(worker *models.Client) error {
	return c.db.Create(worker).Error
}

func (c ClientRepository) Update(worker *models.Client) error {
	return c.db.Updates(worker).Error
}

func (c ClientRepository) Get(id uuid.UUID) (*models.Client, error) {
	client := &models.Client{}
	res := c.db.Where("id = ?", id).First(client)
	return client, res.Error
}

func (c ClientRepository) GetAll() models.Clients {
	var clients models.Clients
	c.db.Find(&clients)
	return clients
}
