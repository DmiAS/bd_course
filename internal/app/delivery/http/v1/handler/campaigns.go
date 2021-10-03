package http

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCampaigns(ctx *gin.Context) {

}

func (h *Handler) updateCampaign(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	req := new(ds.UpdateCampaignInput)
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	camp := converters.ConvertUpdateCampaignInput(req, id)
	if err := h.campaigns.Update(camp); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
