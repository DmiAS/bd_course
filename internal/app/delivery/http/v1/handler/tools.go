package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func extractID(ctx echo.Context) (uuid.UUID, error) {
	sid := ctx.Param("id")
	return uuid.Parse(sid)
}

//func (c *campReq) bind(ctx echo.Context) error {
//	if id := ctx.Param("targetolog_id"); id != "" {
//		var err error
//		c.targetologID, err = uuid.Parse(id)
//		if err != nil {
//			return err
//		}
//	}
//
//	if id := ctx.Param("id"); id != "" {
//		var err error
//		c.id, err = uuid.Parse(id)
//		if err != nil {
//			return err
//		}
//	}
//
//	return ctx.Bind(c)
//}
