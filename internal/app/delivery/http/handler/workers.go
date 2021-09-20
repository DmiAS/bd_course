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

	worker, login := converters.ConvertWorkerInput(req)

	uuid, err := h.services.CreateWorker(worker)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	pass, err := h.services.CreateAuth(uuid, login)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp := converters.ConvertWorkeroutput(login, pass)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateWorker(ctx *gin.Context) {

}

func (h *Handler) getWorkers(ctx *gin.Context) {

}

func (h *Handler) deleteWorker(ctx *gin.Context) {

}
