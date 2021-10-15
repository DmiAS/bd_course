package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type stat struct {
	From string    `json:"from"`
	To   string    `json:"to"`
	ID   uuid.UUID `param:"id"`
}

func (h *Handler) getProjectStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getThreadStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getCampStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getTargetologStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getTargetologCampaignFullStat(ctx echo.Context) error {
	return nil
}
