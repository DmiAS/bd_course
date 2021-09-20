package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
)

func (h *Handler) createClient(ctx *gin.Context) {
	req := new(ds.User)
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	user := converters.ConvertUserInput(req)

	id, err := h.clients.Create(user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.registerUser(id, req.FirstName, req.LastName)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateClient(ctx *gin.Context) {
	req := new(ds.User)

	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uuid")
		return
	}

	client := converters.ConvertUserWitIDInput(req, id)

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
