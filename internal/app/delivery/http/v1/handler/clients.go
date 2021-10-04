package handler

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h *Handler) createClient(ctx echo.Context) error {
	req := &ds.User{}
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	client := converters.ConvertClientInput(req)
	cs := h.cf.GetService(uwork.Admin)
	resp, err := cs.Create(client)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateClient(ctx echo.Context) error {
	req := &ds.User{}
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	client := converters.ConvertClientInputWithID(req, id)
	cs := h.cf.GetService(uwork.Admin)
	if err := cs.Update(client); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getClients(ctx echo.Context) error {
	cs := h.cf.GetService(uwork.Admin)
	clients := cs.GetAll()
	resp := converters.ConvertGetAllClientsOutput(clients)
	return ctx.JSON(http.StatusOK, resp)
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

	resp := converters.ConvertClientOutput(client)

	return ctx.JSON(http.StatusOK, resp)
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
