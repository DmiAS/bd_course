package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createWorker(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	worker := &models.WorkerEntity{}
	if err := ctx.Bind(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.wf.GetService(info.Role)
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
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	worker := &models.WorkerEntity{}
	if err := ctx.Bind(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	ws := h.wf.GetService(info.Role)
	worker.ID = id
	if err := ws.Update(worker); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) getWorkers(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.wf.GetService(info.Role)
	workers := ws.GetAll(pag)
	return ctx.JSON(http.StatusOK, workers)
}

func (h *Handler) getWorker(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	ws := h.wf.GetService(info.Role)
	worker, err := ws.Get(id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, worker)
}

func (h *Handler) getTargetologCampaigns(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ws := h.cmpf.GetService(info.Role)
	camps := ws.GetCampaigns(id, pag)
	return ctx.JSON(http.StatusOK, camps)
}

func (h *Handler) deleteWorker(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	id, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid uuid")
	}

	as := h.af.GetService(info.Role)
	if err := as.Delete(id); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
