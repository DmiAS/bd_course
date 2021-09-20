package ds

import "github.com/google/uuid"

type CreateProjectInput struct {
	Name     string    `json:"name"`
	ClientID uuid.UUID `json:"client_id"`
}

type Project struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetProjectsOutput struct {
	Count    int       `json:"count"`
	Projects []Project `json:"projects"`
}

type CreateProjectOutput struct {
	Project
}
