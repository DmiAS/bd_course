package models

import "github.com/google/uuid"

type Project struct {
	ID       uuid.UUID `json:"id"`
	ClientID uuid.UUID
	Name     string `json:"name"`
}
type Projects []Project

type ProjectsList struct {
	Amount int `json:"amount"`
	Projects
}

func NewProjectsList(ps Projects) *ProjectsList {
	return &ProjectsList{
		Amount:   len(ps),
		Projects: ps,
	}
}
