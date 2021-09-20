package http

import (
	"github.com/gin-gonic/gin"

	"github.com/DmiAS/bd_course/internal/app/service"
)

type Handler struct {
	router  *gin.Engine
	workers service.IWorkerService
	auth    service.IAuthService
}

func NewHandler(workers service.IWorkerService, auth service.IAuthService) *Handler {
	router := gin.Default()
	handler := &Handler{
		router:  router,
		workers: workers,
		auth:    auth,
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
}
