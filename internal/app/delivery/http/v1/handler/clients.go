package handler

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createClient(ctx *gin.Context) {
	req := &ds.User{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	client := converters.ConvertClientInput(req)
	cs := h.cf.GetService(uwork.Admin)
	resp, err := cs.Create(client)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateClient(ctx *gin.Context) {
	req := &ds.User{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	client := converters.ConvertClientInputWithID(req, id)
	cs := h.cf.GetService(uwork.Admin)
	if err := cs.Update(client); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getClients(ctx *gin.Context) {
	cs := h.cf.GetService(uwork.Admin)
	clients := cs.GetAll()
	resp := converters.ConvertGetAllClientsOutput(clients)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getClient(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	cs := h.cf.GetService(uwork.Admin)
	client, err := cs.Get(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters.ConvertClientOutput(client)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) deleteClient(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	as := h.af.GetService(uwork.Admin)
	if err := as.Delete(id); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
