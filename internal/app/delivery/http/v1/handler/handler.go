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
	uf   *service.UserFactory
	pf   *service.ProjectFactory
	tf   *service.ThreadFactory
	cmpf *service.CampaignFactory
	sf   *service.StatsFactory
}

func NewHandler(
	wf *service.WorkerFactory,
	af *service.AuthFactory,
	uf *service.UserFactory,
	pf *service.ProjectFactory,
	tf *service.ThreadFactory,
	cmpf *service.CampaignFactory,
	sf *service.StatsFactory) *Handler {
	router := echo.New()
	handler := &Handler{
		router: router,
		wf:     wf,
		af:     af,
		uf:     uf,
		tf:     tf,
		pf:     pf,
		sf:     sf,
		cmpf:   cmpf,
	}
	handler.initRoutes()
	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(w, req)
}

func (h *Handler) initRoutes() {
	// login handler
	h.router.POST("/login", h.login)
	api := h.router.Group("/api/v1", h.auth)
	{
		admins := api.Group("/admins")
		{
			admins.POST("", h.createAdmin)
			admins.GET("", h.getAdmins)
			admins.GET("/:id", h.getAdmin)
			admins.PUT("/:id", h.updateAdmin)
		}

		workers := api.Group("/workers")
		{
			workers.POST("", h.createWorker)
			workers.GET("", h.getWorkers)
			workers.GET("/:id", h.getWorker)
			workers.PUT("/:id", h.updateWorker)
			workers.DELETE("/:id", h.deleteWorker)
			workers.GET("/:id/camps", h.getTargetologCampaigns)
		}

		auth := api.Group("/auth")
		{
			auth.PUT("/:id", h.updateAuth)
		}

		clients := api.Group("/user")
		{
			clients.POST("/", h.createClient)
			clients.GET("/", h.getClients)
			clients.GET("/:id", h.getClient)
			clients.PUT("/:id", h.updateClient)
			clients.DELETE("/:id", h.deleteClient)
		}

		projects := api.Group("/projects")
		{
			projects.GET("/:client_id", h.getClientProjects)
			projects.POST("/:client_id", h.createProject)
			projects.GET("/:id", h.getProject)
			projects.PUT("/:id", h.updateProject)
			projects.DELETE("/:id", h.deleteProject)
		}

		threads := api.Group("/threads")
		{
			threads.GET("/", h.getThreads)
			threads.POST("/", h.createThread)
			threads.GET("/:id", h.getThread)
			threads.PUT("/:id", h.updateThread)
			threads.DELETE("/:id", h.deleteThread)
		}

		camps := api.Group("/campaigns")
		{
			camps.GET("/", h.getCampaigns)
			camps.PUT("/:id", h.attachCampaign)
			camps.POST("/:id", h.assignCampaign)
		}

		stats := api.Group("/statistic")
		{
			stats.GET("/projects/:id", h.getProjectStat)
			stats.GET("/threads/:thread_id", h.getThreadStat)
			stats.GET("/camps/:camp_id", h.getCampStat)
			stats.GET("/targetologs/:target_id", h.getTargetologStat)
			stats.GET("/targetologs/:target_id/camps/:camp_id", h.getTargetologCampaignFullStat)
		}
	}
}
