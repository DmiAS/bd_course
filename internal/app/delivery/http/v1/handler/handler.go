package http

import (
	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/service"
)

type Handler struct {
	router *gin.Engine

	workers   service.IWorkerService
	auth      service.IAuthService
	projects  service.IProjectService
	threads   service.IThreadService
	campaigns service.ICampaignService
	clients   service.IClientService
}

func NewHandler(workers service.IWorkerService, auth service.IAuthService,
	projects service.IProjectService, threads service.IThreadService,
	campaigns service.ICampaignService,
	clients service.IClientService) *Handler {
	router := gin.Default()
	handler := &Handler{
		router:    router,
		workers:   workers,
		auth:      auth,
		projects:  projects,
		threads:   threads,
		campaigns: campaigns,
	}
	handler.initRoutes()
	return handler
}

func (h *Handler) initRoutes() {
	workers := h.router.Group("/workers")
	{
		workers.POST("/", h.createWorker)
		workers.GET("/", h.getWorkers)
		workers.GET("/:id", h.getWorker)
		workers.PUT("/:id", h.updateWorker)
		workers.DELETE("/:id", h.deleteWorker)
	}

	clients := h.router.Group("/clients")
	{
		clients.POST("/", h.createClient)
		clients.GET("/", h.getClients)
		clients.GET("/", h.getClient)
		clients.PUT("/:id", h.updateClient)
		clients.DELETE("/:id", h.deleteClient)
	}

	projects := h.router.Group("/projects")
	{
		projects.POST("/", h.createProject)
		projects.GET("/", h.getProjects)
		projects.PUT("/:id", h.updateProject)
		projects.DELETE("/:id", h.deleteProject)
	}

	threads := h.router.Group("/threads")
	{
		threads.POST("/", h.createThread)
		threads.GET("/", h.getThreads)
		threads.PUT("/:id", h.updateThread)
		threads.DELETE("/:id", h.deleteThread)
	}

	camps := h.router.Group("/campaigns")
	{
		camps.GET("/", h.getCampaigns)
		camps.PUT("/:id", h.updateCampaign)
	}
}
