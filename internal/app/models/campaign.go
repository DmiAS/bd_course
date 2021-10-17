package models

import (
	"time"

	"github.com/google/uuid"
)

type CampaignStat struct {
	CampID      uuid.UUID
	Date        time.Time
	Spent       float64
	Impressions int
	Conversion  int
	Subs        []int
	Unsubs      []int
	Sales       int
}

type Campaign struct {
	ID           uuid.UUID `json:"id"`
	ThreadID     uuid.UUID `json:"thread_id"`
	TargetologID uuid.UUID `json:"targetolog_id"`
	CabinetID    int       `json:"cabinet_id"`
	ClientID     int       `json:"client_id"`
	VkCampID     int       `json:"vk_camp_id"`
	Name         string    `json:"name"`
}
type Campaigns []Campaign

type CampaignsList struct {
	Campaigns Campaigns
	Amount    int `json:"amount"`
}

func NewCampaignsList(camps Campaigns) *CampaignsList {
	return &CampaignsList{
		Campaigns: camps,
		Amount:    len(camps),
	}
}
