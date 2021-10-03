package http

import (
	converters2 "github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createThread(ctx *gin.Context) {
	req := new(ds.CreateThreadInput)

	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	thread := converters2.ConvertCreateThreadInput(req)
	id, err := h.threads.Create(thread)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters2.ConvertCreateProjectOutput(id, thread.Name)

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) getThreads(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	threads, err := h.threads.Get(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters2.ConvertGetThreadsOutput(threads)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateThread(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	req := new(ds.UpdateThreadInput)
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	thread := converters2.ConvertUpdateThreadInput(req.Name, id)

	if err := h.threads.Update(thread); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) deleteThread(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.threads.Delete(id); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}