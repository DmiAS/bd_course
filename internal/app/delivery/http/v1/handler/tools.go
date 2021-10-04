package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func extractID(ctx echo.Context) (uuid.UUID, error) {
	sid := ctx.Param("id")
	return uuid.Parse(sid)
}

type info struct {
	Name      string `json:"name"`
	clientID  uuid.UUID
	projectID uuid.UUID
}

func (i *info) bind(ctx echo.Context) error {
	clientID := ctx.Param("client_id")
	if clientID != "" {
		id, err := uuid.Parse(clientID)
		if err != nil {
			return err
		}
		i.clientID = id
	}

	projectID := ctx.Param("id")
	if projectID != "" {
		id, err := uuid.Parse(projectID)
		if err != nil {
			return err
		}
		i.projectID = id
	}

	return ctx.Bind(i)
}
