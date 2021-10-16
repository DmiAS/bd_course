package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (h *Handler) login(ctx echo.Context) error {
	info := &Auth{}
	if err := ctx.Bind(info); err != nil {
		return err
	}

	if info.Login == "" || info.Password == "" {
		return ctx.String(http.StatusNonAuthoritativeInfo, "invalid auth data")
	}
	af := h.af.GetService(models.AdminRole)
	token, err := af.Login(info.Login, info.Password)
	if err != nil {
		return ctx.String(http.StatusNonAuthoritativeInfo, "invalid auth data")
	}
	return ctx.JSON(http.StatusOK, tokenResponse{token})
}

func (h *Handler) updateAuth(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	auth := &models.Auth{}
	if err := ctx.Bind(auth); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	auth.UserID = id
	as := h.af.GetService(models.AdminRole)
	if err := as.Update(auth); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
