package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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

	uuid, err := h.services.CreateWorker(worker)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	pass, err := h.services.CreateAuth(uuid, login)
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

	sid := ctx.Param("id")
	id, err := uuid.FromBytes([]byte(sid))
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	worker := converters.ConvertWorkerUpdateInput(req, id)

	if err := h.services.UpdateWorker(worker); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getWorkers(ctx *gin.Context) {

}

func (h *Handler) deleteWorker(ctx *gin.Context) {

}
