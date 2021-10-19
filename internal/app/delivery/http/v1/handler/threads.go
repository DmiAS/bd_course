package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type threadInfo struct {
	Name      string
	ProjectID uuid.UUID
	ThreadID  uuid.UUID
}

func (p *threadInfo) bind(ctx echo.Context) error {
	p.Name = ctx.QueryParam("name")
	tid := ctx.Param("id")
	if tid != "" {
		var err error
		p.ThreadID, err = uuid.Parse(tid)
		if err != nil {
			return err
		}
	}

	pid := ctx.QueryParam("project_id")
	if pid != "" {
		var err error
		p.ProjectID, err = uuid.Parse(pid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) createThread(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	data := &threadInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(info.Role)
	if err := ts.Create(data.ProjectID, data.Name); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusCreated)
}

func (h *Handler) getThread(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	data := &threadInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(info.Role)
	thread, err := ts.Get(data.ThreadID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, thread)
}

func (h *Handler) getThreadCampaigns(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	threadID, err := extractID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	cs := h.cmpf.GetService(info.Role)
	camps := cs.GetThreadCampaigns(threadID, pag)
	return ctx.JSON(http.StatusOK, camps)
}

func (h *Handler) getProjectThreads(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	data := &threadInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(info.Role)
	threads := ts.GetAll(data.ProjectID, pag)
	return ctx.JSON(http.StatusOK, threads)
}

func (h *Handler) updateThread(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	data := &threadInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(info.Role)
	if err := ts.Update(&models.Thread{
		ID:        data.ThreadID,
		ProjectID: data.ProjectID,
		Name:      data.Name,
	}); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) deleteThread(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	data := &threadInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(info.Role)
	if err := ts.Delete(data.ThreadID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
