package handler

import (
	"fmt"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const (
	tokenHeader   = "Authorization"
	tokenTemplate = "Bearer: %s"
	Role          = "role"
	ID            = "id"
)

func (h *Handler) auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenStr, err := getTokenFromHeader(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusNonAuthoritativeInfo, err.Error())
		}
		ids, err := h.af.GetService(uwork.Admin).GetRoleInfo(tokenStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusNonAuthoritativeInfo, err.Error())
		}
		ctx.Set(ID, ids.ID)
		ctx.Set(Role, ids.Role)
		return next(ctx)
	}
}

func getTokenFromHeader(ctx echo.Context) (string, error) {
	var token string
	header := ctx.Request().Header
	authHeader := header.Get(tokenHeader)
	n, err := fmt.Sscanf(authHeader, tokenTemplate, &token)
	if err != nil || n == 0 {
		err = errors.New("access denied")
		return "", err
	}
	return token, nil
}
