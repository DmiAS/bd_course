package campaign

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
)

func (s Service) Get(id uuid.UUID) (*models.Campaign, error) {
	rep := s.unit.GetCampaignsRepository()
	return rep.GetCampaign(id)
}

func (s Service) GetAll() *models.CampaignsList {
	rep := s.unit.GetCampaignsRepository()
	camps := rep.GetAll()
	return models.NewCampaignsList(camps)
}

func (s *Service) GetCampaigns(workerID uuid.UUID) *models.CampaignsList {
	rep := s.unit.GetCampaignsRepository()
	camps := rep.GetCampaigns(workerID)
	return models.NewCampaignsList(camps)
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
