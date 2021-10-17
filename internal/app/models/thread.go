package models

import "github.com/google/uuid"

type Thread struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	Name      string    `json:"name"`
	Created   int64     `json:"-" gorm:"autoCreateTime:nano"`
}
type Threads []Thread

type ThreadsList struct {
	Cursor  int64 `json:"cursor"`
	Amount  int   `json:"amount"`
	Threads `json:"threads"`
}
