package handler

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) updateAuth(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	auth := &ds.Auth{}
	if err := ctx.BindJSON(auth); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	as := h.af.GetService(uwork.Admin)
	if err := as.Update(&models.Auth{
		Login:    auth.Login,
		Password: auth.Password,
		UserID:   id,
	}); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
}
