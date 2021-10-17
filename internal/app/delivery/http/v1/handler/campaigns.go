package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type campInfo struct {
	ThreadID     uuid.UUID
	TargetologID uuid.UUID
	CampaignID   uuid.UUID
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

// получение всех возможных кампаний
func (h *Handler) getCampaigns(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.cmpf.GetService(info.Role)
	camps, err := ws.GetAll(pag)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, camps)
}

func (h *Handler) getCampaign(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	cs := h.cmpf.GetService(info.Role)
	camp, err := cs.Get(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, camp)
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
	if err := ws.AttachCampaign(data.ThreadID, data.CampaignID); err != nil {
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
	if err := ws.AssignCampaign(data.TargetologID, data.CampaignID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
