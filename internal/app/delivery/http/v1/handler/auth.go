package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) login(ctx echo.Context) error {
	info := &Auth{}
	if err := ctx.Bind(info); err != nil {
		return err
	}

	if info.Login == "" || info.Password == "" {
		return ctx.String(http.StatusNonAuthoritativeInfo, "invalid auth data")
	}
	af := h.af.GetService(uwork.Admin)
	token, err := af.Login(info.Login, info.Password)
	if err != nil {
		return ctx.String(http.StatusNonAuthoritativeInfo, "invalid auth data")
	}
	return ctx.JSON(http.StatusOK, token)
}

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
