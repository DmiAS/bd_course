package models

import "github.com/google/uuid"

type Thread struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Name      string
}
type Threads []Thread

type ThreadsList struct {
	Amount int `json:"amount"`
	Threads
}

func NewThreadsList(ps Threads) *ThreadsList {
	return &ThreadsList{
		Amount:  len(ps),
		Threads: ps,
	}
}
