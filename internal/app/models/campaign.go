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
	ID           uuid.UUID
	ThreadID     uuid.UUID
	TargetologID uuid.UUID
	CabinetID    int
	ClientID     int
	VkCampID     int
	Name         string
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
