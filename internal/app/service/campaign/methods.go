package campaign

import (
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) Get(id uuid.UUID) (*models.Campaign, error) {
	rep := s.unit.GetCampaignsRepository()
	return rep.GetCampaign(id)
}

func (s Service) GetAll(cursor int64, limit int) (*models.CampaignsList, error) {
	rep := s.unit.GetCampaignsRepository()
	created := cursor
	if cursor == 0 {
		created = time.Now().UnixNano()
	}
	camps := rep.GetAll(limit, created)
	return createCampaignsList(camps), nil
}

func (s *Service) GetCampaigns(workerID uuid.UUID) *models.CampaignsList {
	rep := s.unit.GetCampaignsRepository()
	camps := rep.GetCampaigns(workerID)
	return createCampaignsList(camps)
}

func (s *Service) AttachCampaign(threadID, campID uuid.UUID) error {
	camp := &models.Campaign{
		ID:       campID,
		ThreadID: threadID,
	}
	rep := s.unit.GetCampaignsRepository()
	return rep.Update(camp)
}

func (s Service) AssignCampaign(targetologID, campID uuid.UUID) error {
	camp := &models.Campaign{
		ID:           campID,
		TargetologID: targetologID,
	}
	rep := s.unit.GetCampaignsRepository()
	return rep.Update(camp)
}

func createCampaignsList(camps models.Campaigns) *models.CampaignsList {
	var cursor int64
	if len(camps)-1 >= 0 {
		cursor = camps[len(camps)-1].Created
	}
	return &models.CampaignsList{
		Cursor:    cursor,
		Campaigns: camps,
		Amount:    len(camps),
	}
}
