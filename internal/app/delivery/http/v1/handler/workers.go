package handler

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h *Handler) createWorker(ctx echo.Context) error {
	req := &ds.Worker{}
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	worker := converters.ConvertWorkerInput(req)
	ws := h.wf.GetService(uwork.Admin)
	resp, err := ws.Create(worker)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateWorker(ctx echo.Context) error {
	req := &ds.Worker{}
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	worker := converters.ConvertUpdateWorkerInput(req, id)
	ws := h.wf.GetService(uwork.Admin)
	if err := ws.Update(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getWorkers(ctx echo.Context) error {
	ws := h.wf.GetService(uwork.Admin)
	workers := ws.GetAll()
	resp := converters.ConvertGetAllWorkerOutput(workers)
	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getWorker(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	ws := h.wf.GetService(uwork.Admin)
	worker, err := ws.Get(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp := converters.ConvertWorkerOutput(worker)

	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) deleteWorker(ctx echo.Context) error {
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
