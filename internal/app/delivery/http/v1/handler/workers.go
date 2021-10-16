package handler

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h *Handler) createWorker(ctx echo.Context) error {
	worker := &models.WorkerEntity{}
	if err := ctx.Bind(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.wf.GetService(models.AdminRole)
	resp, err := ws.Create(worker)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &models.LogPass{
		Login:    resp.Login,
		Password: resp.Password,
	})
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

	ws := h.wf.GetService(models.AdminRole)
	worker.ID = id
	if err := ws.Update(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getWorkers(ctx echo.Context) error {
	ws := h.wf.GetService(models.AdminRole)
	workers := ws.GetAll()
	return ctx.JSON(http.StatusOK, workers)
}

func (h *Handler) getWorker(ctx echo.Context) error {
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	ws := h.wf.GetService(models.AdminRole)
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

	as := h.af.GetService(models.AdminRole)
	if err := as.Delete(id); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
