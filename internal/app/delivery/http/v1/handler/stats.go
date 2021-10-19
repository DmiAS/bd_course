package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type statisticRange struct {
	From         time.Time
	To           time.Time
	CampaignID   uuid.UUID
	ThreadID     uuid.UUID
	TargetologID uuid.UUID
	ProjectID    uuid.UUID
}

func (s *statisticRange) bind(ctx echo.Context) error {
	s.From, s.To = time.Now(), time.Now()
	if from := ctx.QueryParam("from"); from != "" {
		var err error
		s.From, err = time.Parse(models.TimeTemplate, from)
		if err != nil {
			return err
		}
	}

	if to := ctx.QueryParam("to"); to != "" {
		var err error
		s.To, err = time.Parse(models.TimeTemplate, to)
		if err != nil {
			return err
		}
	}

	if campID := ctx.Param("camp_id"); campID != "" {
		var err error
		s.CampaignID, err = uuid.Parse(campID)
		if err != nil {
			return err
		}
	}

	if threadID := ctx.Param("thread_id"); threadID != "" {
		var err error
		s.ThreadID, err = uuid.Parse(threadID)
		if err != nil {
			return err
		}
	}

	if targetID := ctx.Param("target_id"); targetID != "" {
		var err error
		s.TargetologID, err = uuid.Parse(targetID)
		if err != nil {
			return err
		}
	}

	if projectID := ctx.Param("project_id"); projectID != "" {
		var err error
		s.ProjectID, err = uuid.Parse(projectID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) getProjectStat(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	stat := &statisticRange{}
	if err := stat.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ss := h.sf.GetService(info.Role)
	res, err := ss.GetProjectStat(stat.ProjectID, info.ID, info.Role, stat.From, stat.To)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *Handler) getThreadStat(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	stat := &statisticRange{}
	if err := stat.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ss := h.sf.GetService(info.Role)
	res, err := ss.GetThreadStat(stat.ThreadID, stat.From, stat.To)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *Handler) getCampStat(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	stat := &statisticRange{}
	if err := stat.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ss := h.sf.GetService(info.Role)
	res, err := ss.GetFullCampaignStat(stat.CampaignID, info.ID, info.Role, stat.From, stat.To)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *Handler) getTargetologStat(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	stat := &statisticRange{}
	if err := stat.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if err := canManageAccountData(info.Role, info.ID, stat.TargetologID, models.AdminRole); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ss := h.sf.GetService(info.Role)
	res, err := ss.GetTargetologStat(stat.TargetologID, stat.From, stat.To)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
