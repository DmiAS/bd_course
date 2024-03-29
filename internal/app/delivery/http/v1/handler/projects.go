package handler

import (
	"log"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type projectInfo struct {
	Name      string
	ClientID  uuid.UUID
	ProjectID uuid.UUID
}

func (p *projectInfo) bind(ctx echo.Context) error {
	p.Name = ctx.QueryParam("name")
	cid := ctx.QueryParam("client_id")
	if cid != "" {
		var err error
		p.ClientID, err = uuid.Parse(cid)
		if err != nil {
			return err
		}
	}
	pid := ctx.Param("id")
	if pid != "" {
		var err error
		p.ProjectID, err = uuid.Parse(pid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) createProject(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}

	data := &projectInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	ps := h.pf.GetService(info.Role)
	if err := ps.Create(data.ClientID, data.Name); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusCreated)
}

func (h *Handler) getProject(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	data := &projectInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(info.Role)
	project, err := ps.Get(data.ProjectID, info.ID, info.Role)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, project)
}

func (h *Handler) getClientProjects(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	data := &projectInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	pag := &models.Pagination{}
	if err := ctx.Bind(pag); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if err := canManageAccountData(info.Role, info.ID, data.ClientID, models.AdminRole, models.WorkerRole); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(info.Role)
	projects := ps.GetAll(data.ClientID, pag)
	return ctx.JSON(http.StatusOK, projects)
}

func (h *Handler) updateProject(ctx echo.Context) error {
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	data := &projectInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(info.Role)
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
	info, err := extractUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusNonAuthoritativeInfo)
	}
	data := &projectInfo{}
	if err := data.bind(ctx); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ps := h.pf.GetService(info.Role)
	if err := ps.Delete(data.ProjectID); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
