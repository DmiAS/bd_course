package converters

import (
	"github.com/google/uuid"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
)

func ConvertUpdateCampaignInput(req *ds.UpdateCampaignInput, id uuid.UUID) *models.Campaign {
	return &models.Campaign{
		ID:       id,
		ThreadID: req.ThreadID,
		WorkerID: req.WorkerID,
	}
}
