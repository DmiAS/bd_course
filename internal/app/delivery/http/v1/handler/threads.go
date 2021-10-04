package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createThread(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ts := h.tf.GetService(uwork.Admin)
	if err := ts.Create(&models.Thread{
		Name: data.Name,
	}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (h *Handler) getThread(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(uwork.Admin)
	thread, err := ts.Get(data.rootID, data.id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, thread)
}

func (h *Handler) getThreads(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ts := h.tf.GetService(uwork.Admin)
	threads := ts.GetAll(data.rootID)
	return ctx.JSON(http.StatusOK, threads)
}

func (h *Handler) updateThread(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ts := h.tf.GetService(uwork.Admin)
	if err := ts.Update(&models.Thread{
		ID:        data.id,
		ProjectID: data.rootID,
		Name:      data.Name,
	}); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) deleteThread(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ts := h.tf.GetService(uwork.Admin)
	if err := ts.Delete(data.rootID, data.id); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
