package models

import "github.com/google/uuid"

type Project struct {
	ID       uuid.UUID `json:"id"`
	ClientID uuid.UUID `json:"client_id"`
	Name     string    `json:"name"`
	Created  int64     `json:"-" gorm:"autoCreateTime:nano"`
}
type Projects []Project

type ProjectsList struct {
	Cursor   int64 `json:"cursor"`
	Amount   int   `json:"amount"`
	Projects `json:"projects"`
}
