package handler

import (
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/converters"
	"github.com/DmiAS/bd_course/internal/app/delivery/http/v1/ds"
	"github.com/DmiAS/bd_course/internal/app/uwork"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) getProject(ctx *gin.Context) {
	//id, err := extractID(ctx)
	//if err != nil {
	//	ctx.String(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//ps := h.pf.GetService(uwork.Admin)
	//projects, err := ps.GetAll(id)
	//if err != nil {
	//	ctx.String(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//resp := converters.ConvertPro

	//ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) getProjects(ctx *gin.Context) {
	//ps := h.pf.GetService(uwork.Admin)
	//projects := ps.GetAll()
	//resp := converters.ConvertPro

	//ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) updateProject(ctx *gin.Context) {
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

	project := converters.ConvertUpdateProjectInput(req.Name, id)
	ps := h.pf.GetService(uwork.Admin)
	if err := ps.Update(project); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) deleteProject(ctx *gin.Context) {
	id, err := extractID(ctx)
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
