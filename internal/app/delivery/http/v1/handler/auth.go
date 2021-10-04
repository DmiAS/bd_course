package handler

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h *Handler) updateAuth(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	auth := &ds.Auth{}
	if err := ctx.Bind(auth); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	as := h.af.GetService(uwork.Admin)
	if err := as.Update(&models.Auth{
		Login:    auth.Login,
		Password: auth.Password,
		UserID:   id,
	}); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
