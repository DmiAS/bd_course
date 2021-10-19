package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) login(ctx echo.Context) error {
	info := &models.LogPass{}
	if err := ctx.Bind(info); err != nil {
		return err
	}

	if info.Login == "" || info.Password == "" {
		return ctx.String(http.StatusNonAuthoritativeInfo, "invalid auth data")
	}
	af := h.af.GetService(models.AdminRole)
	tokenResp, err := af.Login(info.Login, info.Password)
	if err != nil {
		return ctx.String(http.StatusNonAuthoritativeInfo, "invalid auth data")
	}
	return ctx.JSON(http.StatusOK, tokenResp)
}

func (h *Handler) updateAuth(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
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
	if err := as.Update(auth, info.ID, info.Role); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
