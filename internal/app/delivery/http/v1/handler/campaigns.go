package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type campReq struct {
	ThreadID     uuid.UUID `json:"thread_id"`
	TargetologID uuid.UUID `param:"targetolog_id"`
	CampaignID   uuid.UUID `param:"id"`
}

// получение кампаний таргетолога
func (h *Handler) getTargetologCampaigns(ctx echo.Context) error {
	campID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.cmpf.GetService(uwork.Admin)
	camps := ws.GetCampaigns(campID)
	return ctx.JSON(http.StatusOK, camps)
}

// получение всех компаний, если компания привязана к конкретному
func (h *Handler) getCampaigns(ctx echo.Context) error {
	ws := h.cmpf.GetService(uwork.Admin)
	camps := ws.GetAll()
	return ctx.JSON(http.StatusOK, camps)
}

// привязка кампании к конкретному потоку
func (h *Handler) attachCampaign(ctx echo.Context) error {
	req := &campReq{}
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.cmpf.GetService(uwork.Admin)
	if err := ws.AttachCampaign(req.ThreadID, req.CampaignID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) assignCampaign(ctx echo.Context) error {
	req := &campReq{}
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.cmpf.GetService(uwork.Admin)
	if err := ws.AssignCampaign(req.ThreadID, req.CampaignID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
