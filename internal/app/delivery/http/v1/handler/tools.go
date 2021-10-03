package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func extractID(ctx *gin.Context) (uuid.UUID, error) {
	sid := ctx.Param("id")
	return uuid.FromBytes([]byte(sid))
}
