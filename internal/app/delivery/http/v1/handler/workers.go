package http

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createWorker(ctx *gin.Context) {
	req := &ds.Worker{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	worker := converters.ConvertWorkerInput(req)
	resp, err := h.workers.Create(worker)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateWorker(ctx *gin.Context) {
	req := &ds.Worker{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	worker := converters.ConvertUpdateWorkerInput(req, id)
	if err := h.workers.Update(worker); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getWorkers(ctx *gin.Context) {
	workers, err := h.workers.GetAll()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	resp := converters.ConvertGetAllWorkerOutput(workers)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getWorker(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	worker, err := h.workers.Get(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters.ConvertWorkerOutput(worker)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) deleteWorker(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	if err := h.workers.Delete(id); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}