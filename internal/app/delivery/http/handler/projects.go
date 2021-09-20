package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
)

func (h *Handler) createProject(ctx *gin.Context) {
	req := new(ds.CreateProjectInput)

	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	project := converters.ConvertCreateProjectInput(req)
	id, err := h.projects.Create(project)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters.ConvertCreateProjectOutput(id, project.Name)

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) getProjects(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	projects, err := h.projects.Get(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := converters.ConvertGetProjectsOutput(projects)

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) deleteProject(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.projects.Delete(id); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
