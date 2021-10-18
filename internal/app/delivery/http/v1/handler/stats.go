package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type statisticRange struct {
	From time.Time
	To   time.Time
}

const timeTemplate = "2006-01-02"

func (s *statisticRange) bind(ctx echo.Context) error {
	if from := ctx.QueryParam("from"); from != "" {
		var err error
		s.From, err = time.Parse(timeTemplate, from)
		if err != nil {
			return err
		}
	}

	if to := ctx.QueryParam("to"); to != "" {
		var err error
		s.To, err = time.Parse(timeTemplate, to)
		if err != nil {
			return err

		}
	}
	return nil
}

func (h *Handler) getProjectStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getThreadStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getCampStat(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	campID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	stat := &statisticRange{}
	if err := ctx.Bind(stat); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ss := h.sf.GetService(info.Role)
	res, err := ss.GetFullCampaignStat(campID, stat.From, stat.To)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *Handler) getTargetologStat(ctx echo.Context) error {
	return nil
}

func (h *Handler) getTargetologCampaignFullStat(ctx echo.Context) error {
	return nil
}
