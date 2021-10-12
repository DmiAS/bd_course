package models

import (
	"time"

	"github.com/google/uuid"
)

type ProjectStat struct {
	ProjectID   uuid.UUID          `json:"project_id"`
	From        time.Time          `json:"from"`
	To          time.Time          `json:"to"`
	Spent       float64            `json:"spent"`
	Impressions int                `json:"impressions"`
	Conversion  int                `json:"conversion"`
	Subs        int                `json:"subs"`
	Unsubs      int                `json:"unsubs"`
	Sales       int                `json:"sales"`
	Threads     []ThreadSimpleStat `json:"threads"`
}

type ThreadSimpleStat struct {
	ThreadID    uuid.UUID `json:"thread_id"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Spent       float64   `json:"spent"`
	Impressions int       `json:"impressions"`
	Conversion  int       `json:"conversion"`
	Subs        int       `json:"subs"`
	Unsubs      int       `json:"unsubs"`
	Sales       int       `json:"sales"`
}

type ThreadStat struct {
	ThreadSimpleStat
	Targetologs []int            `json:"targetologs"`
	Camps       []CampSimpleStat `json:"camps"`
}

type CampSimpleStat struct {
	TargetologID uuid.UUID `json:"targetolog_id"`
	CampID       uuid.UUID `json:"camp_id"`
	CabinetID    uuid.UUID `json:"cabinet_id"`
	VkClientID   uuid.UUID `json:"vk_client_id"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
	Spent        float64   `json:"spent"`
	Impressions  int       `json:"impressions"`
	Conversion   int       `json:"conversion"`
	Subs         int       `json:"subs"`
	Unsubs       int       `json:"unsubs"`
	Sales        int       `json:"sales"`
}

type TargetologStat struct {
	Camps []CampSimpleStat `json:"camps"`
}
