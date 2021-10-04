package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func extractID(ctx echo.Context) (uuid.UUID, error) {
	sid := ctx.Param("id")
	return uuid.Parse(sid)
}
