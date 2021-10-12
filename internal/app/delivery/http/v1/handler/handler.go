package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/DmiAS/bd_course/internal/app/service"
)

type Handler struct {
	router *echo.Echo

	wf   *service.WorkerFactory
	af   *service.AuthFactory
	cf   *service.ClientFactory
	pf   *service.ProjectFactory
	tf   *service.ThreadFactory
	cmpf *service.CampaignFactory
	sf   *service.StatsFactory
	//projects  service.IProjectService
	//threads   service.IThreadService
	//campaigns service.ICampaignService
	//clients   service.IClientService
}

func NewHandler(
	wf *service.WorkerFactory,
	af *service.AuthFactory,
	cf *service.ClientFactory,
	pf *service.ProjectFactory,
	tf *service.ThreadFactory,
	cmpf *service.CampaignFactory,
	sf *service.StatsFactory) *Handler {
	router := echo.New()
	handler := &Handler{
		router: router,
		wf:     wf,
		af:     af,
		cf:     cf,
		tf:     tf,
		pf:     pf,
		cmpf:   cmpf,
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
		workers.GET("/:id/camps", h.getTargetologCampaigns)
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
	}

	projects := clients.Group("/:client_id/projects")
	{
		projects.GET("/", h.getProjects)
		projects.POST("/", h.createProject)
		projects.GET("/:id", h.getProject)
		projects.PUT("/:id", h.updateProject)
		projects.DELETE("/:id", h.deleteProject)
	}

	threads := h.router.Group("/projects/:project_id/threads")
	{
		threads.GET("/", h.getThreads)
		threads.POST("/", h.createThread)
		threads.GET("/:id", h.getThread)
		threads.PUT("/:id", h.updateThread)
		threads.DELETE("/:id", h.deleteThread)
	}

	camps := h.router.Group("/campaigns")
	{
		camps.GET("/", h.getCampaigns)
		camps.PUT("/:id", h.attachCampaign)
		camps.POST("/:id", h.assignCampaign)
	}

	stats := h.router.Group("/statistic")
	{
		stats.GET("/projects/:id", h.getProjectStat)
		stats.GET("/threads/:thread_id", h.getThreadStat)
		stats.GET("/camps/:camp_id", h.getCampStat)
		stats.GET("/targetologs/:target_id", h.getTargetologStat)
	}
}
