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
	Name   string `json:"name"`
	rootID uuid.UUID
	id     uuid.UUID
}

func (i *info) bind(ctx echo.Context) error {
	rootID := ctx.Param("client_id")
	if rootID != "" {
		id, err := uuid.Parse(rootID)
		if err != nil {
			return err
		}
		i.rootID = id
	}

	id := ctx.Param("id")
	if id != "" {
		id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		i.id = id
	}

	return ctx.Bind(i)
}

type campReq struct {
	ThreadID     uuid.UUID `json:"thread_id"`
	targetologID uuid.UUID
	id           uuid.UUID
}

func (c *campReq) bind(ctx echo.Context) error {
	if id := ctx.Param("targetolog_id"); id != "" {
		var err error
		c.targetologID, err = uuid.Parse(id)
		if err != nil {
			return err
		}
	}

	if id := ctx.Param("id"); id != "" {
		var err error
		c.id, err = uuid.Parse(id)
		if err != nil {
			return err
		}
	}

	return ctx.Bind(c)
}
