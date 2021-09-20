package ds

import "github.com/google/uuid"

type CreateThreadInput struct {
	Name      string    `json:"name"`
	ProjectID uuid.UUID `json:"project_id"`
}

type Thread struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetThreadsOutput struct {
	Count   int      `json:"count"`
	Threads []Thread `json:"threads"`
}

type CreateThreadOutput struct {
	Thread
}

type UpdateThreadInput struct {
	Name string `json:"name"`
}
