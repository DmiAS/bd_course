package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createAdmin(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	admin := &models.User{Role: models.AdminRole}
	if err := ctx.Bind(admin); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	us := h.uf.GetService(info.Role)
	auth, err := us.Create(admin)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &models.LogPass{
		Login:    auth.Login,
		Password: auth.Password,
	})
}

func (h *Handler) updateAdmin(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	admin := &models.User{}
	if err := ctx.Bind(admin); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	targetID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}
	if targetID != info.ID {
		return ctx.String(http.StatusBadRequest, "you can't modify strangers account")
	}

	admin.ID = targetID
	us := h.uf.GetService(info.Role)
	if err := us.Update(admin); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getAdmins(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	us := h.uf.GetService(info.Role)
	admins := us.GetAll(models.AdminRole)
	return ctx.JSON(http.StatusOK, admins)
}

func (h *Handler) getAdmin(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	targetID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	us := h.uf.GetService(info.Role)
	admin, err := us.Get(targetID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, admin)
}
