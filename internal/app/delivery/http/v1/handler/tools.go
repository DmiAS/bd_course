package http

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func extractID(ctx *gin.Context) (uuid.UUID, error) {
	sid := ctx.Param("id")
	return uuid.FromBytes([]byte(sid))
}

func (h *Handler) registerUser(id uuid.UUID, firstName, lastName string) (*ds.Auth, error) {
	auth, err := h.auth.Create(id, firstName, lastName)
	if err != nil {
		return nil, err
	}

	resp := converters.ConvertAuthOutput(auth)
	return resp, nil
}
