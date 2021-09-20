package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
)

func (h *Handler) createClient(ctx *gin.Context) {
	//req := new(ds.CreateWorkerInput)
	//if err := ctx.BindJSON(req); err != nil {
	//	ctx.String(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//worker, login := converters.ConvertWorkerCreateInput(req)
	//
	//id, err := h.workers.Create(worker)
	//if err != nil {
	//	ctx.String(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//pass, err := h.auth.Create(id, login)
	//if err != nil {
	//	ctx.String(http.StatusInternalServerError, err.Error())
	//}
	//
	//resp := converters.ConvertWorkerCreateOutput(login, pass)
	//
	//ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateClient(ctx *gin.Context) {
	req := new(ds.UpdateClientInput)

	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	client := converters.ConvertUpdateClientInput(req, id)

	if err := h.clients.Update(client); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getClients(ctx *gin.Context) {
	clients, err := h.clients.GetAll()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	resp := converters.ConvertGetAllClientOutput(clients)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getClient(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	client, err := h.clients.Get(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters.ConvertGetClientOutput(client)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) deleteClient(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	if err := h.clients.Delete(id); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
