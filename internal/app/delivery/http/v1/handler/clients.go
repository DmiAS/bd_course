package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createClient(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	client := &models.User{Role: models.ClientRole}
	if err := ctx.Bind(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	us := h.uf.GetService(info.Role)
	auth, err := us.Create(client)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, &models.LogPass{
		Login:    auth.Login,
		Password: auth.Password,
	})
}

func (h *Handler) updateClient(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	targetID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}
	if err := canManageAccountData(info.Role, info.ID, targetID, models.AdminRole); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	client := &models.User{}
	if err := ctx.Bind(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	client.ID = targetID
	us := h.uf.GetService(info.Role)
	if err := us.Update(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getClients(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	if err := canManageAccountData(info.Role, info.ID, uuid.UUID{}, models.AdminRole, models.WorkerRole); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	us := h.uf.GetService(info.Role)
	admins := us.GetAll(models.ClientRole, pag)
	return ctx.JSON(http.StatusOK, admins)
}

func (h *Handler) getClient(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	targetID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}
	if err := canManageAccountData(info.Role, info.ID, targetID, models.AdminRole, models.WorkerRole); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	us := h.uf.GetService(info.Role)
	admin, err := us.Get(targetID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, admin)
}

func (h *Handler) deleteClient(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	targetID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}
	as := h.af.GetService(info.Role)
	if err := as.Delete(targetID); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
