package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

func (h *Handler) createProject(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Create(&models.Project{
		Name: data.Name,
	}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (h *Handler) getProject(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ps := h.pf.GetService(uwork.Admin)
	project, err := ps.Get(data.clientID, data.projectID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, project)
}

func (h *Handler) getProjects(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ps := h.pf.GetService(uwork.Admin)
	projects := ps.GetAll(data.clientID)
	return ctx.JSON(http.StatusOK, projects)
}

func (h *Handler) updateProject(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Update(&models.Project{
		ID:       data.projectID,
		ClientID: data.clientID,
		Name:     data.Name,
	}); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *Handler) deleteProject(ctx echo.Context) error {
	data := &info{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Delete(data.clientID, data.projectID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
