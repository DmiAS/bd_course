package ds

import "github.com/google/uuid"

type UpdateCampaignInput struct {
	ThreadID uuid.UUID `json:"thread_id"`
	WorkerID uuid.UUID `json:"worker_id"`
}

type GetCampaignsOutput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ClientName  string    `json:"client_name"`
	ProjectName string    `json:"project_name"`
	ThreadName  string    `json:"thread_name"`
}
