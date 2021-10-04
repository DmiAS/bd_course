package handler

import (
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/gin-gonic/gin"
)

type info struct {
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

	return nil
}

func (h *Handler) createProject(ctx *gin.Context) {
	req := new(ds.CreateProjectInput)

	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	project := converters.ConvertCreateProjectInput(req)
	ps := h.pf.GetService(uwork.Admin)
	id, err := ps.Create(project)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters.ConvertCreateProjectOutput(id, project.Name)

	ctx.JSON(http.StatusOK, resp)

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
	//ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getProjects(ctx *gin.Context) {
	//ps := h.pf.GetService(uwork.Admin)
	//projects := ps.GetAll()
	//resp := converters.ConvertPro

	//ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateProject(ctx *gin.Context) {
	id, err := extractID(nil)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	req := new(ds.UpdateProjectInput)
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	project := converters.ConvertUpdateProjectInput(req.Name, id)
	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Update(project); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) deleteProject(ctx *gin.Context) {
	id, err := extractID(nil)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Delete(id); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
