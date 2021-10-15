package handler

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h *Handler) createWorker(ctx echo.Context) error {
	worker := &models.WorkerEntity{}
	if err := ctx.Bind(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.wf.GetService(uwork.Admin)
	resp, err := ws.Create(worker)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateWorker(ctx echo.Context) error {
	worker := &models.WorkerEntity{}
	if err := ctx.Bind(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	ws := h.wf.GetService(uwork.Admin)
	worker.ID = id
	if err := ws.Update(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getWorkers(ctx echo.Context) error {
	ws := h.wf.GetService(uwork.Admin)
	workers := ws.GetAll()
	return ctx.JSON(http.StatusOK, workers)
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
	return ctx.JSON(http.StatusOK, worker)
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
