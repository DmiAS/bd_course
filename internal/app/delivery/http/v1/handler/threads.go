package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type threadInfo struct {
	Name      string    `json:"name"`
	ProjectID uuid.UUID `json:"project_id"`
	ThreadID  uuid.UUID `param:"thread_id"`
}

func (h *Handler) createThread(ctx echo.Context) error {
	data := &threadInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ts := h.tf.GetService(uwork.Admin)
	if err := ts.Create(&models.Thread{
		ProjectID: data.ProjectID,
		Name:      data.Name,
	}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (h *Handler) getThread(ctx echo.Context) error {
	data := &threadInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(uwork.Admin)
	thread, err := ts.Get(data.ThreadID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, thread)
}

func (h *Handler) getThreads(ctx echo.Context) error {
	data := &threadInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(uwork.Admin)
	threads := ts.GetAll(data.ProjectID)
	return ctx.JSON(http.StatusOK, threads)
}

func (h *Handler) updateThread(ctx echo.Context) error {
	data := &threadInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ts := h.tf.GetService(uwork.Admin)
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
	data := &threadInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ts := h.tf.GetService(uwork.Admin)
	if err := ts.Delete(data.ThreadID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
