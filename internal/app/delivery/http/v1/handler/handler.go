package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/DmiAS/bd_course/internal/app/service"
)

type Handler struct {
	router *echo.Echo

	wf *service.WorkerFactory
	af *service.AuthFactory
	cf *service.ClientFactory
	pf *service.ProjectFactory
	//projects  service.IProjectService
	//threads   service.IThreadService
	//campaigns service.ICampaignService
	//clients   service.IClientService
}

func NewHandler(
	wf *service.WorkerFactory,
	af *service.AuthFactory,
	cf *service.ClientFactory,
	pf *service.ProjectFactory) *Handler {
	router := echo.New()
	handler := &Handler{
		router: router,
		wf:     wf,
		af:     af,
		cf:     cf,
		//projects:  projects,
		//threads:   threads,
		//campaigns: campaigns,
	}
	handler.initRoutes()
	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(w, req)
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

	auth := h.router.Group("/auth")
	{
		auth.PUT("/:id", h.updateAuth)
	}

	clients := h.router.Group("/clients")
	{
		clients.POST("/", h.createClient)
		clients.GET("/", h.getClients)
		clients.GET("/:id", h.getClient)
		clients.PUT("/:id", h.updateClient)
		clients.DELETE("/:id", h.deleteClient)

		//projects
		//clients.GET("/:client_id/projects", nil)
		//clients.POST("/:client_id/projects", nil)
		//clients.GET("/:client_id/projects/:id", nil)
		//clients.PUT("/:client_id/projects/:id", nil)
		//clients.DELETE("/:client_id/projects/:id", nil)
	}

	projects := clients.Group("/")
	{
		projects.GET("/:client_id/projects", h.getProject)
		projects.POST("/:client_id/projects", h.createProject)
		projects.GET("/:client_id/projects/:id", nil)
		projects.PUT("/:client_id/projects/:id", nil)
		projects.DELETE("/:client_id/projects/:id", nil)
	}
	//
	//	threads := h.router.Group("/threads")
	//	{
	//		threads.POST("/", h.createThread)
	//		threads.GET("/", h.getThreads)
	//		threads.PUT("/:id", h.updateThread)
	//		threads.DELETE("/:id", h.deleteThread)
	//	}
	//
	//	camps := h.router.Group("/campaigns")
	//	{
	//		camps.GET("/", h.getCampaigns)
	//		camps.PUT("/:id", h.updateCampaign)
	//	}
}
