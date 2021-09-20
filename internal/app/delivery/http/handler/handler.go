package http

import (
	"github.com/DmiAS/bd_course/internal/app/service"
	"github.com/gin-gonic/gin"
)

type IService interface {
	service.IWorkerService
	service.IAuthService
}

type Handler struct {
	router   *gin.Engine
	services IService
}

func NewHandler(services IService) *Handler {
	router := gin.Default()
	handler := &Handler{
		router:   router,
		services: services,
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
