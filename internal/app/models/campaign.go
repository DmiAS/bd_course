package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// go representation of campaign_stats table
type CampaignStat struct {
	CampID      uuid.UUID
	Date        time.Time
	Spent       float64
	Impressions int
	Conversion  int
	Subs        pq.Int64Array `gorm:"type:biginteger[]"`
	Unsubs      pq.Int64Array `gorm:"type:biginteger[]"`
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
	Created      int64     `json:"-" gorm:"autoCreateTime:nano"`
}
type Campaigns []Campaign

type CampaignsList struct {
	Cursor    int64     `json:"cursor"`
	Campaigns Campaigns `json:"campaigns"`
	Amount    int       `json:"amount"`
}
