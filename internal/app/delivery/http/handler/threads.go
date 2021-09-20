package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/delivery/http/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/ds"
)

func (h *Handler) createThread(ctx *gin.Context) {
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

func (h *Handler) getThreads(ctx *gin.Context) {
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

func (h *Handler) updateThread(ctx *gin.Context) {
	id, err := extractID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	req := new(ds.UpdateProjectInput)
	if err := ctx.BindJSON(req); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	project := converters.ConvertUpdateInput(req.Name, id)

	if err := h.projects.Update(project); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) deleteThread(ctx *gin.Context) {
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
