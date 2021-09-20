package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
)

func (h *Handler) createWorker(ctx *gin.Context) {
	req := new(ds.CreateWorkerInput)
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	worker, login := converters.ConvertWorkerCreateInput(req)

	id, err := h.workers.Create(worker)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	pass, err := h.auth.Create(id, login)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp := converters.ConvertWorkerCreateOutput(login, pass)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateWorker(ctx *gin.Context) {
	req := new(ds.UpdateWorkerInput)

	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	worker := converters.ConvertWorkerUpdateInput(req, id)

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

	resp := converters.ConvertWorkerGetAllOutput(workers)
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

	resp := converters.ConvertWorkerGetOutput(worker)

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
