package http

import (
	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/service"
)

type Handler struct {
	router   *gin.Engine
	workers  service.IWorkerService
	auth     service.IAuthService
	projects service.IProjectService
	threads  service.IThreadService
}

func NewHandler(workers service.IWorkerService, auth service.IAuthService,
	projects service.IProjectService, threads service.IThreadService) *Handler {
	router := gin.Default()
	handler := &Handler{
		router:   router,
		workers:  workers,
		auth:     auth,
		projects: projects,
		threads:  threads,
	}
	handler.initRoutes()
	return handler
}

func (h *Handler) initRoutes() {
	workers := h.router.Group("/workers")
	{
		workers.POST("/", h.createWorker)
		workers.GET("/", h.getWorkers)
		workers.PUT("/:id", h.updateWorker)
		workers.DELETE("/:id", h.deleteWorker)
	}

	projects := h.router.Group("/projects")
	{
		projects.POST("/", h.createProject)
		projects.GET("/", h.getProjects)
		projects.DELETE("/:id", h.deleteProject)
	}

	//threads := h.router.Group("/threads")
	//{
	//	//threads.POST()
	//}
}
