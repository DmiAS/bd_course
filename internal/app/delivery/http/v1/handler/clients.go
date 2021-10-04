package handler

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h *Handler) createClient(ctx echo.Context) error {
	client := &models.Client{}
	if err := ctx.Bind(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	cs := h.cf.GetService(uwork.Admin)
	resp, err := cs.Create(client)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateClient(ctx echo.Context) error {
	client := &models.Client{}
	if err := ctx.Bind(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	cs := h.cf.GetService(uwork.Admin)
	client.User.ID = id
	if err := cs.Update(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getClients(ctx echo.Context) error {
	cs := h.cf.GetService(uwork.Admin)
	clients := cs.GetAll()
	return ctx.JSON(http.StatusOK, clients)
}

func (h *Handler) getClient(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	cs := h.cf.GetService(uwork.Admin)
	client, err := cs.Get(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, client)
}

func (h *Handler) deleteClient(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	as := h.af.GetService(uwork.Admin)
	if err := as.Delete(id); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
