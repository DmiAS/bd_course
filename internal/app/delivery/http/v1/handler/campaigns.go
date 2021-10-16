package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type campInfo struct {
	ThreadID     uuid.UUID `json:"thread_id"`
	TargetologID uuid.UUID `param:"targetolog_id"`
	CampaignID   uuid.UUID `param:"id"`
}

func (c *campInfo) bind(ctx echo.Context) error {
	threadID := ctx.QueryParam("thread_id")
	if threadID != "" {
		var err error
		c.ThreadID, err = uuid.Parse(threadID)
		if err != nil {
			return err
		}
	}

	targetologID := ctx.QueryParam("targetolog_id")
	if targetologID != "" {
		var err error
		c.TargetologID, err = uuid.Parse(targetologID)
		if err != nil {
			return err
		}
	}

	campaignID := ctx.Param("id")
	if campaignID != "" {
		var err error
		c.CampaignID, err = uuid.Parse(campaignID)
		if err != nil {
			return err
		}
	}
	return nil
}

// получение кампаний таргетолога
func (h *Handler) getTargetologCampaigns(ctx echo.Context) error {
	data := &campInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	ws := h.cmpf.GetService(info.Role)
	camps := ws.GetCampaigns(data.TargetologID)
	return ctx.JSON(http.StatusOK, camps)
}

// получение всех возможных кампаний
func (h *Handler) getCampaigns(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	ws := h.cmpf.GetService(info.Role)
	camps := ws.GetAll()
	return ctx.JSON(http.StatusOK, camps)
}

// привязка кампании к конкретному потоку
func (h *Handler) attachCampaign(ctx echo.Context) error {
	data := &campInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	ws := h.cmpf.GetService(info.Role)
	if err := ws.AttachCampaign(data.TargetologID, data.CampaignID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) assignCampaign(ctx echo.Context) error {
	data := &campInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	ws := h.cmpf.GetService(info.Role)
	if err := ws.AssignCampaign(data.ThreadID, data.CampaignID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
