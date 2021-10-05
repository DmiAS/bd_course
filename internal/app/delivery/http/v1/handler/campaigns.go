package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"
)

// получение кампаний таргетолога
func (h *Handler) getCampaigns(ctx echo.Context) error {
	req := &campReq{}
	if err := req.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.wf.GetService(uwork.Admin)
	camps := ws.GetCampaigns(req.targetologID)
	return ctx.JSON(http.StatusOK, camps)
}

// привязка кампании к конкретному потоку
func (h *Handler) attachCampaign(ctx echo.Context) error {
	req := &campReq{}
	if err := req.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.wf.GetService(uwork.Admin)
	if err := ws.AttachCampaign(req.ThreadID, req.id); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
