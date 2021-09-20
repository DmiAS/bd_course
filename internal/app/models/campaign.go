package models

import "github.com/google/uuid"

type Campaign struct {
	ID        uuid.UUID
	ThreadID  uuid.UUID
	WorkerID  uuid.UUID
	CabinetID int
	ClientID  int
	VkID      int
	Name      string
}
