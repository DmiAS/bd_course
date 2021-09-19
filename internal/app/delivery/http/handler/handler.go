package http

import "github.com/gin-gonic/gin"

type IEndpoint interface {
}

func NewHandler(endpoint IEndpoint) *gin.Engine {
	router := gin.Default()
	initRoutes(router)
	return router
}

func initRoutes(router *gin.Engine) {

}
