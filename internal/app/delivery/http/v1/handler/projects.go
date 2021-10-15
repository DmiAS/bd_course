package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type projectInfo struct {
	Name      string    `json:"name"`
	ClientID  uuid.UUID `json:"client_id"`
	ProjectID uuid.UUID `param:"project_id"`
}

func (h *Handler) createProject(ctx echo.Context) error {
	data := &projectInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Create(&models.Project{
		ClientID: data.ClientID,
		Name:     data.Name,
	}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (h *Handler) getProject(ctx echo.Context) error {
	data := &projectInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ps := h.pf.GetService(uwork.Admin)
	project, err := ps.Get(data.ProjectID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, project)
}

func (h *Handler) getClientProjects(ctx echo.Context) error {
	data := &projectInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ps := h.pf.GetService(uwork.Admin)
	projects := ps.GetAll(data.ClientID)
	return ctx.JSON(http.StatusOK, projects)
}

func (h *Handler) updateProject(ctx echo.Context) error {
	data := &projectInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Update(&models.Project{
		ID:       data.ProjectID,
		ClientID: data.ClientID,
		Name:     data.Name,
	}); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) deleteProject(ctx echo.Context) error {
	data := &projectInfo{}
	if err := ctx.Bind(data); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Delete(data.ProjectID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
